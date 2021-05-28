package repository

import "github.com/ferkze/cryptocompare/types"

type CryptoRepository interface {
	GetLastPrice(fsyms, tsyms []string) ([]types.LastPriceModel, error)
	GetSymbols() ([]types.LastPriceModel, error)
	BulkUpdateLastPrices(models []types.LastPriceModel) error
}
