package product

import (
	"pheet-gin-backend/models"
)

type ProductService interface {
	GetProducts() ([]*models.Product, error)
	GetProduct(id int) (*models.Product, error)
	GetProductByType(coffType string) ([]*models.Product, error)
	GetUser(username string) (*models.User, error)
	Create(product *models.Product) error
	SignUp(user *models.SignUpReq) error
	Update(product *models.Product) error
	Delete(id int) error
}
