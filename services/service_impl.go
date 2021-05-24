package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ferkze/cryptocompare/repository"
	"github.com/ferkze/cryptocompare/types"
)

type cryptoServiceImpl struct {
	repo repository.CryptoRepository
}

func NewCryptoService(repo repository.CryptoRepository) CryptoService {
	return &cryptoServiceImpl{repo}
}

func (s *cryptoServiceImpl) GetLastSymbolsPrice(fsyms, tsyms []string) (types.PricesResponse, error) {
	response, err := s.requestLastSymbolsPrice(fsyms, tsyms)
	if err != nil {
		return response, err
	}
	return response, nil
}
func (s *cryptoServiceImpl) RefreshLastSymbolsPrices(fsyms, tsyms []string) error {
	return nil
}

func (s *cryptoServiceImpl) requestLastSymbolsPrice(fsyms, tsyms []string) (types.PricesResponse, error) {
	prices := types.PricesResponse{}
	from := strings.Join(fsyms, ",")
	to := strings.Join(tsyms, ",")
	url := fmt.Sprintf("https://min-api.cryptocompare.com/data/pricemultifull?fsyms=%s&tsyms=%s", from, to)
	resp, err := http.Get(url)
	if err != nil {
		return prices, err
	}
	err = json.NewDecoder(resp.Body).Decode(&prices)
	return prices, err
}
