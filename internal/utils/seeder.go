package seeder

import (
	"fmt"

	"github.com/darkphotonKN/ecommerce-server-go/config"
	"github.com/darkphotonKN/ecommerce-server-go/internal/models"
	"github.com/darkphotonKN/ecommerce-server-go/internal/product"
)

func SeedProducts() {

	productRepo := product.NewProductRepository(config.DB)
	productService := product.NewProductService(productRepo)

	defaultProducts := []models.Product{
		{Title: "DoggoBites", Subtitle: "Best cookies for your pupper", ImageURL: "test@test.com"},
		{Title: "Kitty Dream Land", Subtitle: "Your kitten's favorite bed.", ImageURL: "test@test.com"},
		{Title: "Orca", Subtitle: "Which orca wouldn't enjoy this?", ImageURL: "test@test.com"},
		{Title: "Bunny Snacks", Subtitle: "Delicious treats for your rabbit", ImageURL: "bunny@test.com"},
		{Title: "Parrot Paradise", Subtitle: "The ultimate perch for your parrot", ImageURL: "parrot@test.com"},
		{Title: "Hamster Haven", Subtitle: "A cozy home for your hamster", ImageURL: "hamster@test.com"},
		{Title: "Fishy Flakes", Subtitle: "Nutritious food for your fish", ImageURL: "fish@test.com"},
		{Title: "Reptile Retreat", Subtitle: "A perfect hideout for your lizard", ImageURL: "reptile@test.com"},
		{Title: "Horse Hoof Care", Subtitle: "Keep your horse's hooves healthy", ImageURL: "horse@test.com"},
		{Title: "Guinea Pig Palace", Subtitle: "Luxury living for your guinea pig", ImageURL: "guineapig@test.com"},
		{Title: "Ferret Fun Tunnel", Subtitle: "An exciting play tunnel for your ferret", ImageURL: "ferret@test.com"},
		{Title: "Turtle Terrarium", Subtitle: "Ideal habitat for your turtle", ImageURL: "turtle@test.com"},
		{Title: "Birdie Bliss", Subtitle: "Tasty treats for your pet bird", ImageURL: "bird@test.com"},
	}

	for _, product := range defaultProducts {
		err := productService.CreateProductService(&product)
		if err != nil {
			fmt.Printf("Error inserting product %s: %v\n", product.Title, err)
		} else {
			fmt.Printf("Inserted product %s successfully\n", product.Title)
		}
	}
}
