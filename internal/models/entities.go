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
	Title    string    `db:"title" json:"title"`
	Subtitle string    `db:"subtitle" json:"subtitle"`
	ImageURL string    `db:"image_url" json:"image_url"`
	Price    int       `db:"price" json:"price"`
	Rating   *float64  `db:"rating" json:"rating"`
	Weight   int       `db:"weight" json:"weight"`
	Detail   *string   `db:"detail" json:"detail"`
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
