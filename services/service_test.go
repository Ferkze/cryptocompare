package services_test

import (
	"log"
	"testing"

	"github.com/ferkze/cryptocompare/services"
)

func TestCryptoServiceImpl(t *testing.T) {

	s := services.NewCryptoService(nil)
	response, err := s.GetLastSymbolsPrice([]string{"BTC"}, []string{"USD"})
	if err != nil {
		t.Fatalf("Error in service GetLastSymbolsPrice: %v\n", err)
	}
	pair, ok := response.RAW["BTC"]
	if !ok {
		t.Fatalf("Error in first pair(from) responded: %v\n", response)
	}
	value, ok := pair["USD"]
	if !ok {
		t.Fatalf("Error in first pair(to) responded: %v\n", response)
	}
	log.Printf("Value received successfully: %v\n", value)

}
