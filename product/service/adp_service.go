package service

import (
	"pheet-gin-backend/models"
	"pheet-gin-backend/product"
)

type productService struct {
	proRepo product.ProductRepository
}

func NewProductService(proRepo product.ProductRepository) productService {
	return productService{proRepo: proRepo}
}

func (r productService) GetProducts() ([]*models.Product, error) {
	return r.proRepo.FetchAll()
}

func (r productService) GetProduct(id int) (*models.Product, error) {
	return r.proRepo.FetchById(id)
}

func (r productService) GetUser(username string) (*models.User, error) {
	return r.proRepo.FetchUser(username)
}

func (r productService) GetProductByType(coffType string) ([]*models.Product, error) {
	return r.proRepo.FetchByType(coffType)
}

func (r productService) Create(product *models.Product) error {
	return r.proRepo.Create(product)
}

func (r productService) SignUp(user *models.SignUpReq) error {
	return r.proRepo.SignUp(user)
}

func (r productService) Update(product *models.Product) error {
	return r.proRepo.Update(product)
}

func (r productService) Delete(id int) error {
	return r.proRepo.Delete(id)
}
