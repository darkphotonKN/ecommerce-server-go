package product

type Product struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	ImageUrl string `json:"imageUrl"`
}

type ProductId struct {
	ID int `json:"id"`
}

type ProductDetail struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	ImageUrl string `json:"imageUrl"`
	price    int    `json:"price"`
	rating   int    `json:"rating"`
	weight   int    `json:"weight"`
	detail   string `json:"detail"`
}
