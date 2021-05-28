package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ferkze/cryptocompare/handler"
	"github.com/ferkze/cryptocompare/jobs"
	"github.com/ferkze/cryptocompare/repository/mysql"
	"github.com/ferkze/cryptocompare/services"
	"github.com/gorilla/mux"
)
var (
	dsn string
	port string
	fsyms []string
	tsyms []string
)

func init() {
	fstr := os.Getenv("fsyms")
	tstr := os.Getenv("tsyms")

	if fstr == "" || tstr == "" {
		panic("No pairs in env parameters (fsyms, tsyms)")
	}
	fsyms = strings.Split(fstr, ",")
	tsyms = strings.Split(tstr, ",")

	dsn = os.Getenv("DB_DSN")
	if dsn == "" {
		panic("No dsn provided")
	}

	port = os.Getenv("port")
	if port == "" {
		port = "8000"
		log.Printf("Using port %s as default\n", port)
	}
}

func main() {
	repo := mysql.NewCryptoRepository(dsn)
	s := services.NewCryptoService(repo)
	h := handler.NewCryptoHandler(repo,s)
	r := mux.NewRouter()
	r.HandleFunc("/api/prices", h.GetPairPrices)

	j := jobs.NewCryptoJobs(s)
	j.InitializeSymbolPrices(fsyms, tsyms)
	j.RefreshSymbolPricesCron(fsyms, tsyms)

	initServer(r)
}

func initServer(r *mux.Router) {
	log.Printf("API running on endpoint: http://localhost:%s/api/prices\n", port)
	http.ListenAndServe(":"+port, r)
}