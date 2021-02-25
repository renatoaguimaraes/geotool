package main

import (
	"net/http"

	"novitatus.com/geotool/internal/adapters/gql"
)

func main() {
	http.Handle("/graphql", corsMiddleware(gql.New()))
	http.ListenAndServe(":8085", nil)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
		next.ServeHTTP(w, r)
	})
}
