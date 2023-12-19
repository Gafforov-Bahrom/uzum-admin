package admin

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/Gafforov-Bahrom/uzum_admin/internal/dto"
	"github.com/gin-gonic/gin"
)

// DeleteProduct(context.Context, dto.TypeID) error

func (r *Router) checkRole(c *gin.Context) error {
	token := c.GetHeader("Authorization")
	out, err := r.getUserRole(c, token)
	if err != nil {
		return err
	}

	if out.Role != "admin" {
		return err
	}

	return nil
}

func (r *Router) AddProduct(c *gin.Context) {
	var request AddProductRequest

	err := r.checkRole(c)
	if err != nil {
		c.AbortWithError(401, err)
		return
	}

	err = c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	id, err := r.productService.CreateProduct(c, dto.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Count,
		Count:       request.Count,
	})

	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, gin.H{
		"id": id,
	})
}

func (r *Router) GetProduct(c *gin.Context) {

	err := r.checkRole(c)
	if err != nil {
		c.AbortWithError(401, err)
		return
	}

	id, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		c.AbortWithError(400, err)
	}

	product, err := r.productService.GetProduct(c, dto.TypeID(id))
	if errors.Is(err, sql.ErrNoRows) {
		c.AbortWithStatus(404)
		return
	}
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, dtoToProduct(product))

}

func (r *Router) GetProducts(c *gin.Context) {
	err := r.checkRole(c)
	if err != nil {
		c.AbortWithError(401, err)
		return
	}
	temp := &dto.ListProductIn{
		Limit:  100,
		Offset: 0,
		Query:  c.Query("query"),
	}

	products, err := r.productService.List(c, temp)
	if errors.Is(err, sql.ErrNoRows) {
		c.AbortWithStatus(404)
	}
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	res := GetProductsResponse{
		Count:    products.Count,
		Products: listToProduct(products.Items),
	}
	c.JSON(200, res)
}

func (r *Router) UpdateProduct(c *gin.Context) {
	err := r.checkRole(c)
	if err != nil {
		c.AbortWithError(401, err)
		return
	}
	var request UpdateProductRequest
	err = c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	id, err := r.productService.UpdateProduct(c, &dto.Product{
		Id:          dto.TypeID(request.Id),
		Name:        request.Name,
		Price:       request.Price,
		Description: request.Description,
		Count:       request.Count,
	})
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, gin.H{
		"id": id,
	})

}

func (r *Router) DeleteProduct(c *gin.Context) {
	err := r.checkRole(c)
	if err != nil {
		c.AbortWithError(401, err)
	}
	id, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	err = r.productService.DeleteProduct(c, dto.TypeID(id))
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	c.JSON(204, gin.H{})
}
