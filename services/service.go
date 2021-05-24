package services

import (
	"github.com/ferkze/cryptocompare/types"
)

type CryptoService interface {
	GetLastSymbolsPrice(fsyms, tsyms []string) (types.PricesResponse, error)
	RefreshLastSymbolsPrices(fsyms, tsyms []string) error
}
