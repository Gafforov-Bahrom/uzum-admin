package http

import (
	"context"
	"net/http"
	"sync"

	"github.com/Gafforov-Bahrom/uzum_admin/internal/transport/http/admin"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	server *http.Server
	engine *gin.Engine
}

func NewServer(port string, engine *gin.Engine, adminRouter *admin.Router) *Server {
	server := &Server{
		engine: engine,
		server: &http.Server{
			Addr: port,
		},
	}
	server.setShopRoutes(adminRouter)
	return server
}

func (s *Server) Start(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	s.server.Handler = s.engine
	go func() {
		<-ctx.Done()
		logrus.Error("shutting down http server")
		logrus.Error(s.server.Shutdown(ctx))
	}()
	return s.server.ListenAndServe()
}

func (s *Server) setShopRoutes(adminRouter *admin.Router) {
	admin := s.engine.Group("/admin")
	admin.POST("/v1/product", adminRouter.AddProduct)
	admin.GET("/v1/product/:product_id", adminRouter.GetProduct)
	admin.GET("/v1/products", adminRouter.GetProducts)
	admin.DELETE("v1/product/:product_id", adminRouter.DeleteProduct)
	admin.PUT("/v1/product", adminRouter.UpdateProduct)
}
