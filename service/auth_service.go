package service

import (
	"github.com/VulpesFerrilata/go-micro-custom/client/grpc"
	"github.com/VulpesFerrilata/grpc/protoc/auth"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
)

func NewAuthService(opts ...client.Option) auth.AuthService {
	service := micro.NewService(
		micro.Name("boardgame.auth.svc.client"),
		micro.Version("latest"),
		micro.Client(
			grpc.NewClient(
				opts...,
			),
		),
	)

	return auth.NewAuthService("boardgame.auth.svc", service.Client())
}
