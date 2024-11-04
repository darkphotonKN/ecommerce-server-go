package product

import (
	"github.com/darkphotonKN/ecommerce-server-go/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	DB *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (r *ProductRepository) GetProducts() (*[]ProductListResponse, error) {
	query := `SELECT id, title, subtitle, image_url FROM products`

	var products []ProductListResponse

	err := r.DB.Select(&products, query)

	if err != nil {
		return nil, err
	}

	return &products, nil
}

func (r *ProductRepository) GetTrendingProducts() (*[]ProductListResponse, error) {
	query := `SELECT id, title, subtitle, image_url FROM products`

	var products []ProductListResponse

	err := r.DB.Select(&products, query)

	if err != nil {
		return nil, err
	}

	return &products, nil
}

func (r *ProductRepository) GetProductById(id uuid.UUID) (*models.Product, error) {
	query := `SELECT * FROM products WHERE products.id = $1`

	var product models.Product

	err := r.DB.Get(&product, query, id)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) GetProductByTitle(title string) (*models.Product, error) {
	query := `SELECT * FROM products WHERE products.title = $1`

	var product models.Product

	err := r.DB.Get(&product, query, title)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) CreateProduct(product *models.Product) error {
	query := `
	INSERT INTO 
	products (title, subtitle, image_url, price, rating, weight, detail) 
	VALUES (:title, :subtitle, :image_url, :price, :rating, :weight, :detail)
	`

	_, err := r.DB.NamedQuery(query, product)

	if err != nil {
		return err
	}

	return nil
}
