package main

import (
	"github.com/Gafforov-Bahrom/uzum_admin/internal/config"
	"github.com/jmoiron/sqlx"
)

type serviceProvider struct {
	cfg *config.Config
	db  *sqlx.DB
	// repoProducts
}

func NewServiceProvider(cfg *config.Config) *serviceProvider {
	return &serviceProvider{cfg: cfg}
}
