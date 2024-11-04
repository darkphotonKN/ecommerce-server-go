package rating

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RatingRepository struct {
	DB *sqlx.DB
}

func NewRatingRepository(db *sqlx.DB) *RatingRepository {
	return &RatingRepository{
		DB: db,
	}
}

func (r *RatingRepository) CreateRating(productId uuid.UUID, ratingReq RatingRequest) error {

	// add a new rating under this product's id
	query := `
	INSERT INTO ratings (product_id, rating)
	VALUES (:product_id, :rating)
	`

	// temporary struct to hold values
	params := map[string]interface{}{
		"product_id": productId,
		"rating":     ratingReq.Rating,
	}

	_, err := r.DB.NamedQuery(query, params)

	if err != nil {
		return err
	}

	return nil
}
