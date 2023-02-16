package repoupdater

import (
	proto "github.com/sourcegraph/sourcegraph/internal/repoupdater/v1"
)

type RepoUpdaterServiceServer struct {
	Server *Server
	proto.UnimplementedRepoUpdaterServiceServer
}
