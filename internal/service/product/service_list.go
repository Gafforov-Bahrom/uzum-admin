package product

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_admin/internal/dto"
)

func (s *Service) List(ctx context.Context, in *dto.ListProductIn) (*dto.ListProductOut, error) {
	return s.repo.List(ctx, in)
}
