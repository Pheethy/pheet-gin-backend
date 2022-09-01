package product

import "pheet-gin-backend/models"

type ProductRepository interface {
	FetchAll() ([]*models.Product, error)
	FetchById(id int)(*models.Product, error)
	FetchByType(coffType string)([]*models.Product, error)
	FetchUser(username string)(*models.User, error)
	Create(product *models.Product)error
	SignUp(user *models.SignUpReq)error
	Update(product *models.Product) error
	Delete(id int)error
}