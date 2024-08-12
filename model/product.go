package model

type Product struct {
	//`json:` especifica o nome que vai ser retornado
	ID    int     `json:"id_product"`
	Name  string  `json:"product_name"`
	Price float64 `json:"product_price"`
}
