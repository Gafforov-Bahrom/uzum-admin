package product

import (
	"github.com/Gafforov-Bahrom/uzum_admin/internal/dto"
	desc "github.com/Gafforov-Bahrom/uzum_admin/pkg/grpc/admin"
)

func dtoToDesc(in *dto.Product) *desc.Product {
	return &desc.Product{
		Id:          uint64(in.Id),
		Name:        in.Name,
		Description: in.Description,
		Price:       in.Price,
		Count:       in.Count,
	}
}

func listToDesc(in []*dto.Product) []*desc.Product {
	out := make([]*desc.Product, len(in))
	for i, item := range in {
		out[i] = dtoToDesc(item)
	}
	return out
}
