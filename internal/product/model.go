package product

import "github.com/google/uuid"

type ProductListResponse struct {
	ID       uuid.UUID `db:"id" json:"id"`
	Title    string    `db:"title" json:"title"`
	Subtitle string    `db:"subtitle" json:"subtitle"`
	ImageURL string    `db:"image_url" json:"image_url"`
}
