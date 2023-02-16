package repoupdater

import (
	"context"

	"github.com/sourcegraph/sourcegraph/internal/api"
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
