package product

type Product struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	ImageUrl string `json:"imageUrl"`
}
