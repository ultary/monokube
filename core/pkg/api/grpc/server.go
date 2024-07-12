package grpc

import (
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/ultary/monokube/core/pkg/api/grpc/v1"
)

type Server struct {
	server *grpc.Server
}

func NewServer() *Server {

	server := grpc.NewServer()
	v1.RegisterKlusterServer(server)
	v1.RegisterSystemServer(server)

	return &Server{
		server: server,
	}
}

func (s *Server) Serve(network, address string) error {
	lis, err := net.Listen(network, address)
	if err != nil {
		log.Errorf("[gRPC] Failed to listen grpc: %v", err)
		return err
	}
	if err = s.server.Serve(lis); err != nil {
		log.Errorf("[gRPC] Failed to serve grpc: %v", err)
		return err
	}
	return nil
}

func (s *Server) Stop() {
	log.Info("[gRPC] Stopping grpc server")
	s.server.Stop()
	log.Info("[gRPC] Stopped grpc server")
}
