package product

type Product struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	ImageUrl string `json:"imageUrl"`
	Price    int    `json:"price"`
	Rating   int    `json:"rating"`
	Weight   int    `json:"weight"`
	Detail   string `json:"detail"`
}

type ProductSummary struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	ImageUrl string `json:"imageUrl"`
}
