package repoupdater

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sourcegraph/sourcegraph/internal/api"
	"github.com/sourcegraph/sourcegraph/internal/env"
	internalgrpc "github.com/sourcegraph/sourcegraph/internal/grpc"
	"github.com/sourcegraph/sourcegraph/internal/grpc/defaults"
	"github.com/sourcegraph/sourcegraph/internal/httpcli"
	"github.com/sourcegraph/sourcegraph/internal/repoupdater/protocol"
	proto "github.com/sourcegraph/sourcegraph/internal/repoupdater/v1"
	"github.com/sourcegraph/sourcegraph/internal/syncx"
	"github.com/sourcegraph/sourcegraph/internal/trace/ot"
	"github.com/sourcegraph/sourcegraph/lib/errors"
)

var (
	repoUpdaterURL = env.Get("REPO_UPDATER_URL", "http://repo-updater:3182", "repo-updater server URL")

	grpcClient = syncx.OnceValues(func() (proto.RepoUpdaterServiceClient, error) {
		u, err := url.Parse(repoUpdaterURL)
		if err != nil {
			return nil, err
		}
		// TODO: how important is it to use a context here?
		conn, err := grpc.Dial(u.Host, defaults.DialOptions()...)
		if err != nil {
			return nil, err
		}

		return proto.NewRepoUpdaterServiceClient(conn), nil
	})

	defaultDoer, _ = httpcli.NewInternalClientFactory("repoupdater").Doer()

	// DefaultClient is the default Client. Unless overwritten, it is
	// connected to the server specified by the REPO_UPDATER_URL
	// environment variable.
	DefaultClient = NewClient(repoUpdaterURL)
)

// Client is a repoupdater client.
type Client struct {
	// URL to repoupdater server.
	URL string

	// HTTP client to use
	HTTPClient httpcli.Doer
}

// NewClient will initiate a new repoupdater Client with the given serverURL.
func NewClient(serverURL string) *Client {
	return &Client{
		URL:        serverURL,
		HTTPClient: defaultDoer,
	}
}

// RepoUpdateSchedulerInfo returns information about the state of the repo in the update scheduler.
func (c *Client) RepoUpdateSchedulerInfo(
	ctx context.Context,
	args protocol.RepoUpdateSchedulerInfoArgs,
) (result *protocol.RepoUpdateSchedulerInfoResult, err error) {
	if internalgrpc.IsGRPCEnabled(ctx) {
		client, err := grpcClient()
		if err != nil {
			return nil, err
		}
		req := &proto.RepoUpdateSchedulerInfoRequest{Id: int32(args.ID)}
		resp, err := client.RepoUpdateSchedulerInfo(ctx, req)
		if err != nil {
			return nil, err
		}
		var result protocol.RepoUpdateSchedulerInfoResult
		result.FromProto(resp)
		return &result, nil
	}

	resp, err := c.httpPost(ctx, "repo-update-scheduler-info", args)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		stack := fmt.Sprintf("RepoScheduleInfo: %+v", args)
		return nil, errors.Wrap(errors.Errorf("http status %d", resp.StatusCode), stack)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

// MockRepoLookup mocks (*Client).RepoLookup for tests.
var MockRepoLookup func(protocol.RepoLookupArgs) (*protocol.RepoLookupResult, error)

// RepoLookup retrieves information about the repository on repoupdater.
func (c *Client) RepoLookup(
	ctx context.Context,
	args protocol.RepoLookupArgs,
) (result *protocol.RepoLookupResult, err error) {
	if MockRepoLookup != nil {
		return MockRepoLookup(args)
	}

	span, ctx := ot.StartSpanFromContext(ctx, "Client.RepoLookup") //nolint:staticcheck // OT is deprecated
	defer func() {
		if result != nil {
			span.SetTag("found", result.Repo != nil)
		}
		if err != nil {
			ext.Error.Set(span, true)
			span.SetTag("err", err.Error())
		}
		span.Finish()
	}()
	if args.Repo != "" {
		span.SetTag("Repo", string(args.Repo))
	}

	if internalgrpc.IsGRPCEnabled(ctx) {
		client, err := grpcClient()
		if err != nil {
			return nil, err
		}
		resp, err := client.RepoLookup(ctx, args.ToProto())
		if err != nil {
			return nil, errors.Wrapf(err, "RepoLookup for %+v failed", args)
		}
		switch {
		case resp.GetErrorNotFound():
			return nil, &ErrNotFound{Repo: args.Repo, IsNotFound: true}
		case resp.GetErrorUnauthorized():
			return nil, &ErrUnauthorized{Repo: args.Repo, NoAuthz: true}
		case resp.GetErrorTemporarilyUnavailable():
			return nil, &ErrTemporary{Repo: args.Repo, IsTemporary: true}
		}
		return protocol.RepoLookupResultFromProto(resp), nil
	}

	resp, err := c.httpPost(ctx, "repo-lookup", args)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// best-effort inclusion of body in error message
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 200))
		return nil, errors.Errorf(
			"RepoLookup for %+v failed with http status %d: %s",
			args,
			resp.StatusCode,
			string(body),
		)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err == nil && result != nil {
		switch {
		case result.ErrorNotFound:
			err = &ErrNotFound{
				Repo:       args.Repo,
				IsNotFound: true,
			}
		case result.ErrorUnauthorized:
			err = &ErrUnauthorized{
				Repo:    args.Repo,
				NoAuthz: true,
			}
		case result.ErrorTemporarilyUnavailable:
			err = &ErrTemporary{
				Repo:        args.Repo,
				IsTemporary: true,
			}
		}
	}
	return result, err
}

// MockEnqueueRepoUpdate mocks (*Client).EnqueueRepoUpdate for tests.
var MockEnqueueRepoUpdate func(ctx context.Context, repo api.RepoName) (*protocol.RepoUpdateResponse, error)

// EnqueueRepoUpdate requests that the named repository be updated in the near
// future. It does not wait for the update.
func (c *Client) EnqueueRepoUpdate(ctx context.Context, repo api.RepoName) (*protocol.RepoUpdateResponse, error) {
	if MockEnqueueRepoUpdate != nil {
		return MockEnqueueRepoUpdate(ctx, repo)
	}

	if internalgrpc.IsGRPCEnabled(ctx) {
		client, err := grpcClient()
		if err != nil {
			return nil, err
		}

		req := proto.EnqueueRepoUpdateRequest{Repo: string(repo)}
		resp, err := client.EnqueueRepoUpdate(ctx, &req)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				return nil, &repoNotFoundError{repo: string(repo), responseBody: err.Error()}
			}
			return nil, err
		}

		return protocol.RepoUpdateResponseFromProto(resp), nil
	}

	req := &protocol.RepoUpdateRequest{
		Repo: repo,
	}

	resp, err := c.httpPost(ctx, "enqueue-repo-update", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	var res protocol.RepoUpdateResponse
	if resp.StatusCode == http.StatusNotFound {
		return nil, &repoNotFoundError{string(repo), string(bs)}
	} else if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return nil, errors.New(string(bs))
	} else if err = json.Unmarshal(bs, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type repoNotFoundError struct {
	repo         string
	responseBody string
}

func (repoNotFoundError) NotFound() bool { return true }
func (e *repoNotFoundError) Error() string {
	return fmt.Sprintf("repo %v not found with response: %v", e.repo, e.responseBody)
}

// MockEnqueueChangesetSync mocks (*Client).EnqueueChangesetSync for tests.
var MockEnqueueChangesetSync func(ctx context.Context, ids []int64) error

func (c *Client) EnqueueChangesetSync(ctx context.Context, ids []int64) error {
	if MockEnqueueChangesetSync != nil {
		return MockEnqueueChangesetSync(ctx, ids)
	}

	if internalgrpc.IsGRPCEnabled(ctx) {
		client, err := grpcClient()
		if err != nil {
			return err
		}

		// empty response can be ignored
		_, err = client.EnqueueChangesetSync(ctx, &proto.EnqueueChangesetSyncRequest{Ids: ids})
		return err
	}

	req := protocol.ChangesetSyncRequest{IDs: ids}
	resp, err := c.httpPost(ctx, "enqueue-changeset-sync", req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read response body")
	}

	var res protocol.ChangesetSyncResponse
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return errors.New(string(bs))
	} else if err = json.Unmarshal(bs, &res); err != nil {
		return err
	}

	if res.Error == "" {
		return nil
	}
	return errors.New(res.Error)
}

// MockSchedulePermsSync mocks (*Client).SchedulePermsSync for tests.
var MockSchedulePermsSync func(ctx context.Context, args protocol.PermsSyncRequest) error

func (c *Client) SchedulePermsSync(ctx context.Context, args protocol.PermsSyncRequest) error {
	if MockSchedulePermsSync != nil {
		return MockSchedulePermsSync(ctx, args)
	}

	if internalgrpc.IsGRPCEnabled(ctx) {
		client, err := grpcClient()
		if err != nil {
			return err
		}

		_, err = client.SchedulePermsSync(ctx, args.ToProto())
		return err
	}

	resp, err := c.httpPost(ctx, "schedule-perms-sync", args)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "read response body")
	}

	var res protocol.PermsSyncResponse
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return errors.New(string(bs))
	} else if err = json.Unmarshal(bs, &res); err != nil {
		return err
	}

	if res.Error == "" {
		return nil
	}
	return errors.New(res.Error)
}

// MockSyncExternalService mocks (*Client).SyncExternalService for tests.
var MockSyncExternalService func(ctx context.Context, externalServiceID int64) (*protocol.ExternalServiceSyncResult, error)

// SyncExternalService requests the given external service to be synced.
func (c *Client) SyncExternalService(ctx context.Context, externalServiceID int64) (*protocol.ExternalServiceSyncResult, error) {
	if MockSyncExternalService != nil {
		return MockSyncExternalService(ctx, externalServiceID)
	}

	if internalgrpc.IsGRPCEnabled(ctx) {
		client, err := grpcClient()
		if err != nil {
			return nil, err
		}

		_, err = client.SyncExternalService(ctx, &proto.SyncExternalServiceRequest{ExternalServiceId: externalServiceID})
		return nil, err
	}

	req := &protocol.ExternalServiceSyncRequest{ExternalServiceID: externalServiceID}
	resp, err := c.httpPost(ctx, "sync-external-service", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	var result protocol.ExternalServiceSyncResult
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return nil, errors.New(string(bs))
	} else if len(bs) == 0 {
		return &result, nil
	} else if err = json.Unmarshal(bs, &result); err != nil {
		return nil, err
	}

	if result.Error != "" {
		return nil, errors.New(result.Error)
	}
	return &result, nil
}

// MockExternalServiceNamespaces mocks (*Client).QueryExternalServiceNamespaces for tests.
var MockExternalServiceNamespaces func(ctx context.Context, args protocol.ExternalServiceNamespacesArgs) (*protocol.ExternalServiceNamespacesResult, error)

// ExternalServiceNamespaces retrieves a list of namespaces available to the given external service configuration
func (c *Client) ExternalServiceNamespaces(ctx context.Context, args protocol.ExternalServiceNamespacesArgs) (result *protocol.ExternalServiceNamespacesResult, err error) {
	if MockExternalServiceNamespaces != nil {
		return MockExternalServiceNamespaces(ctx, args)
	}

	resp, err := c.httpPost(ctx, "external-service-namespaces", args)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err == nil && result != nil && result.Error != "" {
		err = errors.New(result.Error)
	}
	return result, err
}

func (c *Client) httpPost(ctx context.Context, method string, payload any) (resp *http.Response, err error) {
	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.URL+"/"+method, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	return c.do(ctx, req)
}

func (c *Client) do(ctx context.Context, req *http.Request) (_ *http.Response, err error) {
	span, ctx := ot.StartSpanFromContext(ctx, "Client.do") //nolint:staticcheck // OT is deprecated
	defer func() {
		if err != nil {
			ext.Error.Set(span, true)
			span.SetTag("err", err.Error())
		}
		span.Finish()
	}()

	req.Header.Set("Content-Type", "application/json")

	req = req.WithContext(ctx)
	req, ht := nethttp.TraceRequest(span.Tracer(), req,
		nethttp.OperationName("RepoUpdater Client"),
		nethttp.ClientTrace(false))
	defer ht.Finish()

	if c.HTTPClient != nil {
		return c.HTTPClient.Do(req)
	}
	return http.DefaultClient.Do(req)
}
