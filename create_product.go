package main

import (
	"encoding/json"
	"net/http"
)


func creatProduct(w http.ResponseWriter, r *http.Request) {
	
	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid product data", 400)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	sendData(w, newProduct, 201)
}
