package service

import (
	"github.com/VulpesFerrilata/grpc/protoc/auth"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
)

func NewAuthService(opts ...client.Option) auth.AuthService {
	service := micro.NewService(
		micro.Name("boardgame.auth.svc.client"),
		micro.Version("latest"),
		micro.Client(
			client.NewClient(
				opts...,
			),
		),
	)

	return auth.NewAuthService("boardgame.auth.svc", service.Client())
}
