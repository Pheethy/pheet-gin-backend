package product

import (
	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	GetProducts(c *gin.Context)
	GetProductById(c *gin.Context)
}