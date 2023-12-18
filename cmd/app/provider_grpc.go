package main

import (
	"github.com/Gafforov-Bahrom/uzum_admin/internal/transport/grpc/product"
	desc "github.com/Gafforov-Bahrom/uzum_admin/pkg/grpc/admin"
)

func (sp *serviceProvider) GetProductServer() desc.AdminServiceServer {
	return product.NewService(
		sp.GetProductService(),
	)
}
