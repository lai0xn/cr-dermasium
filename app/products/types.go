package products

type ProductPayload struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	InStock     int    `json:"inStock"`
	Picture     string `json:"image"`
}
