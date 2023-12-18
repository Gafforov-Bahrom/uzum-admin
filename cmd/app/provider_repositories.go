package main

import "github.com/Gafforov-Bahrom/uzum_admin/internal/repository/product"

func (sp *serviceProvider) GetProductRepository() *product.Repository {
	if sp.repoProducts == nil {
		sp.repoProducts = product.NewRepository(
			sp.GetDB(),
		)
	}

	return sp.repoProducts
}
