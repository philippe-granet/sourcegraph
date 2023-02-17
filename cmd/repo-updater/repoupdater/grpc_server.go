package repoupdater

import (
	"context"
	"net/http"

	"github.com/sourcegraph/sourcegraph/internal/api"
	"github.com/sourcegraph/sourcegraph/internal/authz"
	"github.com/sourcegraph/sourcegraph/internal/repoupdater/protocol"
	proto "github.com/sourcegraph/sourcegraph/internal/repoupdater/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RepoUpdaterServiceServer struct {
	Server *Server
	proto.UnimplementedRepoUpdaterServiceServer
}

func (s *RepoUpdaterServiceServer) RepoUpdateSchedulerInfo(ctx context.Context, req *proto.RepoUpdateSchedulerInfoRequest) (*proto.RepoUpdateSchedulerInfoResponse, error) {
	res := s.Server.Scheduler.ScheduleInfo(api.RepoID(req.GetId()))
	return res.ToProto(), nil
}

func (s *RepoUpdaterServiceServer) RepoLookup(ctx context.Context, req *proto.RepoLookupRequest) (*proto.RepoLookupResponse, error) {
	args := protocol.RepoLookupArgs{
		Repo:   api.RepoName(req.Repo),
		Update: req.Update,
	}
	res, err := s.Server.repoLookup(ctx, args)
	if err != nil {
		return nil, err
	}
	return res.ToProto(), nil
}

func (s *RepoUpdaterServiceServer) EnqueueRepoUpdate(ctx context.Context, req *proto.EnqueueRepoUpdateRequest) (*proto.EnqueueRepoUpdateResponse, error) {
	args := &protocol.RepoUpdateRequest{
		Repo: api.RepoName(req.GetRepo()),
	}
	res, httpStatus, err := s.Server.enqueueRepoUpdate(ctx, args)
	if err != nil {
		if httpStatus == http.StatusNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	return &proto.EnqueueRepoUpdateResponse{
		Id:   int32(res.ID),
		Name: res.Name,
	}, nil
}

func (s *RepoUpdaterServiceServer) EnqueueChangesetSync(ctx context.Context, req *proto.EnqueueChangesetSyncRequest) (*proto.EnqueueChangesetSyncResponse, error) {
	if s.Server.ChangesetSyncRegistry == nil {
		s.Server.Logger.Warn("ChangesetSyncer is nil")
		return nil, status.Error(codes.Internal, "changeset syncer is not configured")
	}

	if len(req.Ids) == 0 {
		return nil, status.Error(codes.InvalidArgument, "no ids provided")
	}

	return &proto.EnqueueChangesetSyncResponse{}, s.Server.ChangesetSyncRegistry.EnqueueChangesetSyncs(ctx, req.Ids)
}

func (s *RepoUpdaterServiceServer) SchedulePermsSync(ctx context.Context, req *proto.SchedulePermsSyncRequest) (*proto.SchedulePermsSyncResponse, error) {
	if s.Server.DatabaseBackedPermissionSyncerEnabled != nil && s.Server.DatabaseBackedPermissionSyncerEnabled(ctx) {
		s.Server.Logger.Warn("Dropping schedule-perms-sync request because PermissionSyncWorker is enabled. This should not happen.")
		return &proto.SchedulePermsSyncResponse{}, nil
	}

	if s.Server.PermsSyncer == nil {
		return nil, status.Error(codes.Internal, "perms syncer not configured")
	}

	repoIDs := make([]api.RepoID, len(req.GetRepoIds()))
	for i, id := range req.GetRepoIds() {
		repoIDs[i] = api.RepoID(id)
	}

	if len(req.UserIds) == 0 && len(repoIDs) == 0 {
		return nil, status.Error(codes.InvalidArgument, "neither user IDs nor repo IDs was provided in request (must provide at least one)")
	}

	opts := authz.FetchPermsOptions{InvalidateCaches: req.GetOptions().GetInvalidateCaches()}
	s.Server.PermsSyncer.ScheduleUsers(ctx, opts, req.UserIds...)
	s.Server.PermsSyncer.ScheduleRepos(ctx, repoIDs...)

	return &proto.SchedulePermsSyncResponse{}, nil
}
