package socket

import (
	context "context"

	"github.com/samekigor/feather-deamon/internal/proto"
)

func (s *Server) Run(ctx context.Context, in *proto.RunRequest) (*proto.RunResponse, error) {
	return nil, nil
}

func (s *Server) Start(ctx context.Context, in *proto.StartRequest) (*proto.StartResponse, error) {
	return nil, nil
}

func (s *Server) Stop(ctx context.Context, in *proto.StopRequest) (*proto.StopResponse, error) {
	return nil, nil
}

func (s *Server) Restart(ctx context.Context, in *proto.RestartRequest) (*proto.RestartResponse, error) {
	return nil, nil
}

func (s *Server) Status(ctx context.Context, in *proto.StatusRequest) (*proto.StatusResponse, error) {
	return nil, nil
}

func (s *Server) Delete(ctx context.Context, in *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	return nil, nil
}
