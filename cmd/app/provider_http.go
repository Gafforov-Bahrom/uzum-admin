package main

import (
	"github.com/Gafforov-Bahrom/uzum_admin/internal/transport/http"
	"github.com/Gafforov-Bahrom/uzum_admin/internal/transport/http/admin"
	"github.com/gin-gonic/gin"
)

func (sp *serviceProvider) GetAdminRouter() *admin.Router {
	return admin.NewRouter(
		sp.GetProductService(),
		sp.GetLoginClient(),
	)
}

func (sp *serviceProvider) GetHTTPServer() *http.Server {
	return http.NewServer(sp.cfg.HTTPPort, gin.Default(), sp.GetAdminRouter())
}
