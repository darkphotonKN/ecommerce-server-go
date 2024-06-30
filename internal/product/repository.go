package product

var products = []Product{
	{ID: 1, Title: "DoggoBites", Subtitle: "Best biscuits for your pupper", ImageUrl: "test@test.com"},
	{ID: 2, Title: "Kitty Dream Land", Subtitle: "Your kitten's favorite bed.", ImageUrl: "test@test.com"},
	{ID: 3, Title: "Orca", Subtitle: "Which orca wouldn't enjoy this?", ImageUrl: "test@test.com"},
}

func GetAllProducts() []Product {
	return products
}

func GetAllTrendingProducts() []Product {
	return products
}
