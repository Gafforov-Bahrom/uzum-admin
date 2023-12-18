package main

import (
	"github.com/Gafforov-Bahrom/uzum_admin/internal/service/product"
	desc "github.com/Gafforov-Bahrom/uzum_admin/pkg/grpc/login_v1"
	"google.golang.org/grpc"
)

func (sp *serviceProvider) GetProductService() *product.Service {
	if sp.serviceProducts == nil {
		sp.serviceProducts = product.NewService(
			sp.GetProductRepository(),
		)
	}
	return sp.serviceProducts
}

func (sp *serviceProvider) GetLoginClient() desc.LoginV1Client {
	connection, err := grpc.Dial("localhost:9081", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return desc.NewLoginV1Client(connection)
}
