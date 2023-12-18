package admin

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_admin/internal/dto"
)

type ProductService interface {
	List(context.Context, *dto.ListProductIn) (*dto.ListProductOut, error)
	GetProduct(context.Context, dto.TypeID) (*dto.Product, error)
}
