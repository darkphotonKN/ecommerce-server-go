package product

var products = []Product{
	{ID: 1, Title: "DoggoBites", Subtitle: "Best cookies for your pupper", ImageUrl: "test@test.com", Detail: "A tasty treat for your best doggo friend, look deep into the warm fuzzy dough that makes up these cookies.", Price: 120, Rating: 4, Weight: 107},
	{ID: 2, Title: "Kitty Dream Land", Subtitle: "Your kitten's favorite bed.", ImageUrl: "test@test.com",
		Detail: "Kitties were born for this dreamland.", Price: 2000, Rating: 5, Weight: 207,
	},
	{ID: 3, Title: "Orca", Subtitle: "Which orca wouldn't enjoy this?", ImageUrl: "test@test.com",
		Detail: "Have you ever wondered what your baby orca would feel if they had their own private aquatic bed? Legendary!", Price: 60000, Rating: 1, Weight: 8297,
	},
	{ID: 4, Title: "Bunny Snacks", Subtitle: "Delicious treats for your rabbit", ImageUrl: "bunny@test.com", Detail: "Crunchy and healthy snacks that will make your bunny hop with joy.", Price: 50, Rating: 5, Weight: 200},
	{ID: 5, Title: "Parrot Paradise", Subtitle: "The ultimate perch for your parrot", ImageUrl: "parrot@test.com", Detail: "A beautifully crafted perch that gives your parrot a perfect spot to relax.", Price: 80, Rating: 4, Weight: 300},
	{ID: 6, Title: "Hamster Haven", Subtitle: "A cozy home for your hamster", ImageUrl: "hamster@test.com", Detail: "A spacious and secure habitat designed to keep your hamster happy and safe.", Price: 70, Rating: 5, Weight: 1500},
	{ID: 7, Title: "Fishy Flakes", Subtitle: "Nutritious food for your fish", ImageUrl: "fish@test.com", Detail: "High-quality fish flakes that provide all the essential nutrients your fish needs.", Price: 25, Rating: 4, Weight: 100},
	{ID: 8, Title: "Reptile Retreat", Subtitle: "A perfect hideout for your lizard", ImageUrl: "reptile@test.com", Detail: "A natural-looking hideout that offers a safe and comfortable place for your reptile.", Price: 90, Rating: 5, Weight: 1200},
	{ID: 9, Title: "Horse Hoof Care", Subtitle: "Keep your horse's hooves healthy", ImageUrl: "horse@test.com", Detail: "A premium hoof conditioner that strengthens and protects your horse's hooves.", Price: 150, Rating: 5, Weight: 500},
	{ID: 10, Title: "Guinea Pig Palace", Subtitle: "Luxury living for your guinea pig", ImageUrl: "guineapig@test.com", Detail: "A spacious and comfortable cage that provides a perfect living environment for your guinea pig.", Price: 110, Rating: 4, Weight: 2500},
	{ID: 11, Title: "Ferret Fun Tunnel", Subtitle: "An exciting play tunnel for your ferret", ImageUrl: "ferret@test.com", Detail: "A colorful and durable tunnel that provides endless fun for your ferret.", Price: 60, Rating: 4, Weight: 300},
	{ID: 12, Title: "Turtle Terrarium", Subtitle: "Ideal habitat for your turtle", ImageUrl: "turtle@test.com", Detail: "A well-designed terrarium that offers the perfect environment for your turtle.", Price: 200, Rating: 5, Weight: 3500},
	{ID: 13, Title: "Birdie Bliss", Subtitle: "Tasty treats for your pet bird", ImageUrl: "bird@test.com", Detail: "Delicious and nutritious treats that will make your bird chirp with happiness.", Price: 40, Rating: 4, Weight: 250},
}

func GetAllProducts() []Product {
	return products
}

func GetAllTrendingProducts() []Product {
	return products
}
