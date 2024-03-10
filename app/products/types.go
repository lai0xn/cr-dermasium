package products

type ProductPayload struct {
	Name        string `json:"name"        schema:"name,required"`
	Price       int    `json:"price"       schema:"price,required"`
	Description string `json:"description" schema:"description,required"`
	InStock     int    `json:"inStock"     schema:"inStock,required"`
	Picture     string `json:"image"`
}
