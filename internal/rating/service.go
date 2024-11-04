package rating

import (
	"github.com/google/uuid"
)

type RatingService struct {
	Repo *RatingRepository
}

func NewRatingService(repo *RatingRepository) *RatingService {
	return &RatingService{
		Repo: repo,
	}
}

func (s *RatingService) PostRatingService(productId uuid.UUID, ratingReq RatingRequest) error {
	return s.Repo.CreateRating(productId, ratingReq)
}
