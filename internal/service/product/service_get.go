package product

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_admin/internal/dto"
)

func (s *Service) GetProduct(ctx context.Context, id dto.TypeID) (*dto.Product, error) {
	return s.repo.GetProduct(ctx, id)
}
