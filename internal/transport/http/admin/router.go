package admin

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_admin/pkg/grpc/login_v1"
)

type Router struct {
	productService ProductService
	loginClient    login_v1.LoginV1Client
}

func NewRouter(productService ProductService, loginClient login_v1.LoginV1Client) *Router {
	return &Router{
		productService: productService,
		loginClient:    loginClient,
	}
}

func (s *Router) getUserId(ctx context.Context, token string) (uint64, error) {
	out, err := s.loginClient.GetUserId(ctx, &login_v1.GetUserIdRequest{
		AccessToken: token,
	})
	if err != nil {
		return 0, err
	}
	return out.UserId, nil
}
