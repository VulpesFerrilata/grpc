package gateway

import (
	"github.com/VulpesFerrilata/grpc/protoc/auth"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
)

type AuthGateway interface {
	auth.AuthService
}

func NewAuthGateway(opts ...client.Option) AuthGateway {
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
