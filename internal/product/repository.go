package product

import (
	"github.com/darkphotonKN/ecommerce-server-go/internal/models"
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

func (r *ProductRepository) GetProducts() (*[]models.Product, error) {
	query := `SELECT * FROM products`

	var products []models.Product

	err := r.DB.Select(&products, query)

	if err != nil {
		return nil, err
	}

	return &products, nil
}

func (r *ProductRepository) GetTrendingProducts() (*[]models.Product, error) {
	query := `SELECT * FROM products`

	var products []models.Product

	err := r.DB.Select(&products, query)

	if err != nil {
		return nil, err
	}

	return &products, nil
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
