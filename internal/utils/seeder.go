package seeder

import (
	"fmt"

	"github.com/darkphotonKN/ecommerce-server-go/config"
	"github.com/darkphotonKN/ecommerce-server-go/internal/models"
	"github.com/darkphotonKN/ecommerce-server-go/internal/product"
	"github.com/google/uuid"
)

func SeedProducts() {
	productRepo := product.NewProductRepository(config.DB)
	productService := product.NewProductService(productRepo)
	id, _ := uuid.Parse("adac23a1-a288-45d9-8be4-27597107d6c2")
	idTwo, _ := uuid.Parse("276f6ac4-c350-41fc-8389-7949bc1e28ea")
	idThree, _ := uuid.Parse("9ad35806-c260-4016-9898-e87af6126326")
	idFour, _ := uuid.Parse("526ef978-6b90-4174-bc62-f83e30f95a04")
	idFive, _ := uuid.Parse("2feb2331-d783-4665-a543-6be5bf20de9d")
	idSix, _ := uuid.Parse("d31c98cc-f23d-47a0-909d-96bf286802a6")
	idSeven, _ := uuid.Parse("2249cad1-4568-4dc8-8a12-41be593f25fd")
	idEight, _ := uuid.Parse("78910178-cec4-4e1d-9687-1750beb7caf2")
	idNine, _ := uuid.Parse("1f5c9637-290c-4554-a619-42904dd43e42")
	idTen, _ := uuid.Parse("e011cd86-8d37-4931-ba54-b3b76a364be3")
	idEleven, _ := uuid.Parse("c953ebf1-8cae-41ac-8e18-bdf76b985cab")
	idTwelve, _ := uuid.Parse("dadc7798-9ffb-42a1-b2c6-d7c8ce028cba")

	defaultProducts := []models.Product{
		{ID: id, Title: "DoggoBites", Subtitle: "Best cookies for your pupper", ImageURL: "test@test.com"},
		{ID: idTwo, Title: "Kitty Dream Land", Subtitle: "Your kitten's favorite bed.", ImageURL: "test@test.com"},
		{ID: idThree, Title: "Orca", Subtitle: "Which orca wouldn't enjoy this?", ImageURL: "test@test.com"},
		{ID: idFour, Title: "Bunny Snacks", Subtitle: "Delicious treats for your rabbit", ImageURL: "bunny@test.com"},
		{ID: idFive, Title: "Parrot Paradise", Subtitle: "The ultimate perch for your parrot", ImageURL: "parrot@test.com"},
		{ID: idSix, Title: "Hamster Haven", Subtitle: "A cozy home for your hamster", ImageURL: "hamster@test.com"},
		{ID: idSeven, Title: "Fishy Flakes", Subtitle: "Nutritious food for your fish", ImageURL: "fish@test.com"},
		{ID: idEight, Title: "Reptile Retreat", Subtitle: "A perfect hideout for your lizard", ImageURL: "reptile@test.com"},
		{ID: idNine, Title: "Horse Hoof Care", Subtitle: "Keep your horse's hooves healthy", ImageURL: "horse@test.com"},
		{ID: idTen, Title: "Guinea Pig Palace", Subtitle: "Luxury living for your guinea pig", ImageURL: "guineapig@test.com"},
		{ID: idEleven, Title: "Ferret Fun Tunnel", Subtitle: "An exciting play tunnel for your ferret", ImageURL: "ferret@test.com"},
		{ID: idTwelve, Title: "Turtle Terrarium", Subtitle: "Ideal habitat for your turtle", ImageURL: "turtle@test.com"},
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
