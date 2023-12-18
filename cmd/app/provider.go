package main

import (
	"github.com/Gafforov-Bahrom/uzum_admin/internal/config"
	repoProduct "github.com/Gafforov-Bahrom/uzum_admin/internal/repository/product"
	serviceProducts "github.com/Gafforov-Bahrom/uzum_admin/internal/service/product"
	"github.com/jmoiron/sqlx"
)

type serviceProvider struct {
	cfg             *config.Config
	db              *sqlx.DB
	repoProducts    *repoProduct.Repository
	serviceProducts *serviceProducts.Service
}

func NewServiceProvider(cfg *config.Config) *serviceProvider {
	return &serviceProvider{cfg: cfg}
}
