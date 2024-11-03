package product

import (
	"github.com/darkphotonKN/ecommerce-server-go/internal/models"
	"github.com/google/uuid"
)

type ProductService struct {
	Repo *ProductRepository
}

func NewProductService(repo *ProductRepository) *ProductService {
	return &ProductService{
		Repo: repo,
	}
}

func (s *ProductService) GetProductsService() (*[]ProductListResponse, error) {
	return s.Repo.GetProducts()
}

func (s *ProductService) GetProductById(id uuid.UUID) (*models.Product, error) {
	return s.Repo.GetProductById(id)
}

func (s *ProductService) CreateProductService(product *models.Product) error {
	return s.Repo.CreateProduct(product)
}

func (s *ProductService) GetTrendingProductsService() (*[]ProductListResponse, error) {
	return s.Repo.GetTrendingProducts()
}
