package service

import (
	"basic/internal/config"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func StartGrpcService(config *config.Config) (*grpc.Server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Api.Host, config.Api.Port))
	if err != nil {
		return nil, err
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	// pb.RegisterPatientServiceServer(grpcServer, api.GetBasicGrpcApiServiceServer(*config))

	grpcServer.Serve(lis)

	return grpcServer, nil
}
