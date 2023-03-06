package main

type Product struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	StockLevel  int    `json:"stockLevel"`
}

var stock = []Product{
	{
		Id:          1,
		Description: "Jacket",
		StockLevel:  0,
	},
	{
		Id:          2,
		Description: "Shoes",
		StockLevel:  0,
	},
}

func findProduct(id int) (match Product, found bool) {
	for _, product := range stock {
		if product.Id == id {
			return product, true
		}
	}
	return match, false
}
