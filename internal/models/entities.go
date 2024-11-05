package models

import (
	"github.com/google/uuid"
	"time"
)

/**
* Shared entities that are imported by more than one package.
**/
type User struct {
	BaseDBDateModel
	Email    string `db:"email" json:"email"`
	Name     string `db:"name" json:"name"`
	Password string `db:"password" json:"password,omitempty"`
}

type Product struct {
	ID       uuid.UUID `db:"id" json:"id"`
	Title    string    `db:"title" json:"title" validate:"required"`
	Subtitle string    `db:"subtitle" json:"subtitle" validate:"required"`
	ImageURL string    `db:"image_url" json:"image_url" validate:"required,url"`
	Price    int       `db:"price" json:"price" validate:"required,min=1"`
	Rating   *float64  `db:"rating" json:"rating" validate:"omitempty,gte=0,lte=5"`
	Weight   int       `db:"weight" json:"weight" validate:"required"`
	Detail   *string   `db:"detail" json:"detail"`
}

type Rating struct {
	ID        uuid.UUID `db:"id" json:"id"`
	ProductID uuid.UUID `db:"product_id" json:"productId"`
	Rating    int       `db:"rating" json:"rating"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

/**
* Base models for default table columns.
**/
type BaseDBUserModel struct {
	ID          uuid.UUID `db:"id" json:"id"`
	UpdatedUser uuid.UUID `db:"updated_user" json:"updatedUser"`
	CreatedUser uuid.UUID `db:"created_user" json:"createdUser"`
}

type BaseDBUserDateModel struct {
	ID          uuid.UUID `db:"id" json:"id"`
	UpdatedUser uuid.UUID `db:"updated_user" json:"updatedUser"`
	CreatedUser uuid.UUID `db:"created_user" json:"createdUser"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type BaseDBDateModel struct {
	ID        uuid.UUID `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
