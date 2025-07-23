package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != "GET" {
		http.Error(w, "Please use GET method", 400)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/products", getProducts)
	fmt.Println("Server is running on port 8080...")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
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
	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)
	productList = append(productList, prd6)
}
