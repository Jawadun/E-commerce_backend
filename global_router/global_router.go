package global_router

import "net/http"

func Globalroute(mux *http.ServeMux) http.Handler {
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,DELETE,PATCH")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type Habib")

		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		mux.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handleAllReq)

}
