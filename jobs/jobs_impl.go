package jobs

import (
	"log"

	"github.com/ferkze/cryptocompare/services"
	cron "github.com/robfig/cron/v3"
)


type cryptoPricesJobsImpl struct {
	s services.CryptoService
	c *cron.Cron
}

func NewCryptoJobs(s services.CryptoService) CryptoPricesJobs {
	c := cron.New()
	return &cryptoPricesJobsImpl{s, c}
}

func (j *cryptoPricesJobsImpl) InitializeSymbolPrices(fsyms, tsyms []string) error {
	err := j.s.RefreshLastSymbolsPrices(fsyms, tsyms)
	if err != nil {
		log.Printf("Error executing job InitializeSymbolPrices(): %s\n", err.Error())
		return err
	}
	log.Println("Job InitializeSymbolPrices successfully executed!")
	return nil
}

func (j *cryptoPricesJobsImpl) RefreshSymbolPricesCron(fsyms, tsyms []string) error {
	// Updates each minute
	_, err := j.c.AddFunc("0 * * * * *", func() {
		err := j.s.RefreshLastSymbolsPrices(fsyms, tsyms)
		if err != nil {
			log.Printf("Error executing job RefreshSymbolPricesCron(): %s\n", err.Error())
			return
		}
		log.Println("Job RefreshSymbolPricesCron successfully executed!")
	})
	return err
}
