package product

import (
	"fmt"

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

func (s *ProductService) GetProductsService(limit, offset int) (*[]ProductListResponse, error) {
	return s.Repo.GetProducts(limit, offset)
}

func (s *ProductService) GetProductById(id uuid.UUID) (*models.Product, error) {
	// acquire product first
	product, err := s.Repo.GetProductById(id)

	if err != nil {
		return nil, err
	}

	// acquire its ratings and average them
	ratings, err := s.RatingService.GetAllRatingsForProductService(id)

	fmt.Printf("ratings for product %s: %+v\n", id, ratings)

	if err != nil {
		return nil, err
	}

	avgRating := s.AverageRatings(ratings)

	// update product

	product.Rating = avgRating

	return product, nil
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

// --- Helpers ---

func (s *ProductService) AverageRatings(ratings *[]models.Rating) *float64 {
	if len(*ratings) == 0 {
		return nil
	}

	var totalRating float64 = 0

	for _, rating := range *ratings {
		totalRating += float64(rating.Rating)
	}

	avg := totalRating / float64(len(*ratings))

	return &avg
}
