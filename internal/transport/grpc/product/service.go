package product

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_admin/internal/dto"
	desc "github.com/Gafforov-Bahrom/uzum_admin/pkg/grpc/admin"
	"github.com/Gafforov-Bahrom/uzum_admin/pkg/grpc/login_v1"
)

type grpcService struct {
	desc.UnimplementedAdminServiceServer
	productService ProductService
	loginClient    login_v1.LoginV1Client
}

func NewService(productService ProductService, loginClient login_v1.LoginV1Client) desc.AdminServiceServer {
	return &grpcService{
		productService: productService,
		loginClient:    loginClient,
	}
}

func (s *grpcService) getUserRole(ctx context.Context, token string) (*dto.GetUserRoleOut, error) {
	out, err := s.loginClient.GetUserRole(ctx, &login_v1.GetUserRoleRequest{
		AccessToken: token,
	})
	if err != nil {
		return nil, err
	}

	return &dto.GetUserRoleOut{
		UserId: out.UserId,
		Role:   out.Role,
	}, nil
}
