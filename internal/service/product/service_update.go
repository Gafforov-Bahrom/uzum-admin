package product

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_admin/internal/dto"
)

func (s *Service) UpdateProduct(ctx context.Context, in *dto.Product) (dto.TypeID, error) {
	return s.repo.UpdateProduct(ctx, in)
}
