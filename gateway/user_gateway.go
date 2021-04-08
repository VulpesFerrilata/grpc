package gateway

import (
	"github.com/VulpesFerrilata/grpc/protoc/user"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
)

type UserGateway interface {
	user.UserService
}

func NewUserGateway(opts ...client.Option) UserGateway {
	service := micro.NewService(
		micro.Name("boardgame.user.svc.client"),
		micro.Version("latest"),
		micro.Client(
			client.NewClient(
				opts...,
			),
		),
	)

	return user.NewUserService("boardgame.user.svc", service.Client())
}
