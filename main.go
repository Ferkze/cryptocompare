package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ferkze/cryptocompare/handler"
	"github.com/ferkze/cryptocompare/repository/mysql"
	"github.com/ferkze/cryptocompare/services"
	"github.com/gorilla/mux"
)


func main() {
	repo := mysql.NewCryptoRepository(os.Getenv("DB_DSN"))
	s := services.NewCryptoService(repo)
	
	h := handler.NewCryptoHandler(repo,s)

	r := mux.NewRouter()

	r.HandleFunc("/api/prices", h.GetPairPrices)

	port := os.Getenv("port")
	if port == "" {
		port = "8000"
		log.Printf("Using port %s as default\n", port)
	}
	log.Printf("API running on endpoint: http://localhost:%s/api/prices\n", port)
	http.ListenAndServe(":"+port, r)
}
