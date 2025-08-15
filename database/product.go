package database

var ProductList []Product

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

func init() {
	prd1 := Product{
		ID:          1,
		Name:        "Laptop",
		Description: "A high performance laptop",
		Price:       999.99,
		ImgURL:      "http://example.com/laptop.jpg",
	}
	prd2 := Product{
		ID:          2,
		Name:        "Smartphone",
		Description: "A latest model smartphone",
		Price:       699.99,
		ImgURL:      "http://example.com/smartphone.jpg",
	}
	prd3 := Product{
		ID:          3,
		Name:        "Headphones",
		Description: "Noise-cancelling headphones",
		Price:       199.99,
		ImgURL:      "http://example.com/headphones.jpg",
	}
	prd4 := Product{
		ID:          4,
		Name:        "Smartwatch",
		Description: "A smartwatch with fitness tracking",
		Price:       249.99,
		ImgURL:      "http://example.com/smartwatch.jpg",
	}
	prd5 := Product{
		ID:          5,
		Name:        "Tablet",
		Description: "A versatile tablet for work and play",
		Price:       499.99,
		ImgURL:      "http://example.com/tablet.jpg",
	}
	prd6 := Product{
		ID:          6,
		Name:        "Bluetooth Speaker",
		Description: "Portable Bluetooth speaker with great sound",
		Price:       89.99,
		ImgURL:      "http://example.com/speaker.jpg",
	}
	ProductList = append(ProductList, prd1, prd2, prd3, prd4, prd5, prd6)
}
