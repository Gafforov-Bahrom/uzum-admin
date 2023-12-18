package product

import desc "github.com/Gafforov-Bahrom/uzum_admin/pkg/grpc/admin"

type grpcService struct {
	desc.UnimplementedAdminServiceServer
	productService ProductService
}

func NewService(productService ProductService) desc.AdminServiceServer {
	return &grpcService{
		productService: productService,
	}
}
