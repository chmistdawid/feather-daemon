package socket

import (
	context "context"
	"log"

	"github.com/samekigor/feather-deamon/internal/proto"
)

func (s *Server) Run(ctx context.Context, in *proto.RunRequest) (*proto.RunResponse, error) {
	log.Printf("Run request received: %v", in)
	return nil, nil
}

func (s *Server) Start(ctx context.Context, in *proto.StartRequest) (*proto.StartResponse, error) {
	log.Printf("Start request received: %v", in)
	return nil, nil
}

func (s *Server) Stop(ctx context.Context, in *proto.StopRequest) (*proto.StopResponse, error) {
	log.Printf("Stop request received: %v", in)
	return nil, nil
}

func (s *Server) Restart(ctx context.Context, in *proto.RestartRequest) (*proto.RestartResponse, error) {
	log.Printf("Restart request received: %v", in)
	return nil, nil
}

func (s *Server) Status(ctx context.Context, in *proto.StatusRequest) (*proto.StatusResponse, error) {
	log.Printf("Status request received: %v", in)
	return nil, nil
}

func (s *Server) Delete(ctx context.Context, in *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	log.Printf("Delete request received: %v", in)
	return nil, nil
}
