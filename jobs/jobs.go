package jobs

type CryptoPricesJobs interface {
	InitializeSymbolPrices(fsyms, tsyms []string) error

	RefreshSymbolPricesCron(fsyms, tsyms []string) error
}
