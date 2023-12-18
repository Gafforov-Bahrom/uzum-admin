package product

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_admin/internal/dto"
)

func (s *Service) DeleteProduct(ctx context.Context, in dto.TypeID) error {
	return s.repo.DeleteProduct(ctx, in)
}
