package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	globalrouter := global_router.Globalroute(mux)

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreatProduct))
	mux.Handle("GET /products/{productId}", http.HandlerFunc(handlers.GetProductByID))

	fmt.Println("Server is running on port 8080...")

	err := http.ListenAndServe(":8080", globalrouter) // Use globalrouter to handle all requests
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
