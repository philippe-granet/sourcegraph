package repoupdater

import (
	"context"

	"github.com/sourcegraph/sourcegraph/internal/api"
	"github.com/sourcegraph/sourcegraph/internal/repoupdater/protocol"
	proto "github.com/sourcegraph/sourcegraph/internal/repoupdater/v1"
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
