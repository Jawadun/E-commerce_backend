package handlers

import (
	"encoding/json"
	"net/http"
	"ecommerce/database"
	"ecommerce/utils"
)

func CreatProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid product data", 400)
		return
	}

	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)

	utils.SendData(w, newProduct, 201)
}
