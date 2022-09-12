package route

import (
	"pheet-gin-backend/product"

	"github.com/gin-gonic/gin"
)

type Route struct {
	e *gin.RouterGroup
}

func NewRoute(e *gin.RouterGroup) Route {
	return Route{e: e}
}

func (r Route) RegisterProduct(handler product.ProductHandler) {
	r.e.GET("/products", handler.GetProducts)
	r.e.GET("/product/:id", handler.GetProductById)
}
