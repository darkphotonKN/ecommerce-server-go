package product

import (
	"github.com/darkphotonKN/ecommerce-server-go/internal/models"
	"github.com/darkphotonKN/ecommerce-server-go/internal/rating"
	"github.com/google/uuid"
)

type ProductService struct {
	Repo          *ProductRepository
	RatingService *rating.RatingService
}

func NewProductService(repo *ProductRepository, ratingService *rating.RatingService) *ProductService {
	return &ProductService{
		Repo:          repo,
		RatingService: ratingService,
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

func (s *ProductService) PostRatingService(productId uuid.UUID, ratingReq rating.RatingRequest) error {
	return s.RatingService.PostRatingService(productId, ratingReq)
}
