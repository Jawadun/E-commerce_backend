package handlers

import (
	"ecommerce/utils"
	"ecommerce/database"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("productId")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Invalid product ID", 400)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == pId {
			utils.SendData(w, product, 200)
			return
		}
	}
	http.Error(w, "Product not found", 404)
}
