package product

import "github.com/darkphotonKN/ecommerce-server-go/internal/models"

type ProductService struct {
	Repo *ProductRepository
}

func NewProductService(repo *ProductRepository) *ProductService {
	return &ProductService{
		Repo: repo,
	}
}

func (s *ProductService) GetProductsService() (*[]models.Product, error) {
	return s.Repo.GetProducts()
}

func (s *ProductService) CreateProductService(product *models.Product) error {
	return s.Repo.CreateProduct(product)
}

func (s *ProductService) GetTrendingProductsService() (*[]models.Product, error) {
	return s.Repo.GetTrendingProducts()
}
