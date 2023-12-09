package api

import (
	"basic/internal/config"
)

type basicGrpcApi struct {
	config config.Config
}

func GetBasicGrpcApiServiceServer(config config.Config) *basicGrpcApi {
	return &basicGrpcApi{
		config: config,
	}
}
