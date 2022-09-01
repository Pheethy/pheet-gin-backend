package handler

import (
	"net/http"
	"os"
	"pheet-gin-backend/product"
	"pheet-gin-backend/models"
	"pheet-gin-backend/auth"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type productHandler struct {
	proServ product.ProductService
}

func NewProductHandler(proServ product.ProductService) productHandler {
	return productHandler{proServ: proServ}
}

func (p productHandler) GetProducts(c *gin.Context) {
	products, err := p.proServ.GetProducts()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	resp := map[string]interface{}{
		"products": products,
	}

	c.JSON(http.StatusOK, resp)
}

func (h productHandler) Login(c *gin.Context) {
	var request = models.User{}

	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	if request.UserName == "" || request.Password == "" {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	user, err := h.proServ.GetUser(request.UserName)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(request.Password))
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	tokenz := auth.AccessToken(os.Getenv("SIGN"))

	resp := map[string]interface{}{
		"message": "Login-success",
		"jwt": tokenz,
	}

	c.JSON(http.StatusOK, resp)
}
	
