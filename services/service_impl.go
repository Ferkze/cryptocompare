package services

import (
	"encoding/json"
	"errors"
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

func (s *cryptoServiceImpl) GetLastSymbolsPrice(fsyms, tsyms []string) (prices types.PricesResponse, err error) {
	prices, err = s.requestLastSymbolsPrice(fsyms, tsyms)
	if err == nil {
		return prices, err
	}
	priceRows, err := s.repo.GetLastPrice(fsyms, tsyms)
	if err != nil {
		return prices, err
	}
	if len(priceRows) != 0 {
		return s.transformCryptoModel(priceRows), err
	}
	return prices, errors.New("price data could not be retrieved")
}

func (s *cryptoServiceImpl) RefreshLastSymbolsPrices(fsyms, tsyms []string) error {
	prices, err := s.requestLastSymbolsPrice(fsyms, tsyms)
	if err != nil {
		return err
	}
	models := s.getLastPriceModels(prices)
	return s.repo.BulkUpdateLastPrices(models)
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

func (s *cryptoServiceImpl) transformCryptoModel(models []types.LastPriceModel) types.PricesResponse {
	prices := types.PricesResponse{}
	for _, model := range models {
		rawpair, ok := prices.RAW[model.FROMSYMBOLDISPLAY]
		if !ok {
			rawpair = make(map[string]types.PriceRaw)
		}
		rawpair[model.TOSYMBOLDISPLAY] = types.PriceRaw{
			CHANGE24HOUR:    model.CHANGE24HOUR,
			CHANGEPCT24HOUR: model.CHANGEPCT24HOUR,
			OPEN24HOUR:      model.OPEN24HOUR,
			VOLUME24HOUR:    model.VOLUME24HOUR,
			VOLUME24HOURTO:  model.VOLUME24HOURTO,
			LOW24HOUR:       model.LOW24HOUR,
			HIGH24HOUR:      model.HIGH24HOUR,
			PRICE:           model.PRICE,
			LASTUPDATE:      model.LASTUPDATE,
			SUPPLY:          model.SUPPLY,
			MKTCAP:          model.MKTCAP,
		}
		displaypair, ok := prices.DISPLAY[model.FROMSYMBOLDISPLAY]
		if !ok {
			displaypair = make(map[string]types.PriceDisplay)
		}
		displaypair[model.TOSYMBOLDISPLAY] = types.PriceDisplay{
			CHANGE24HOUR:    fmt.Sprintf("%s %.2f", model.TOSYMBOLDISPLAY, model.CHANGE24HOUR),
			CHANGEPCT24HOUR: fmt.Sprintf("%.2f", model.CHANGEPCT24HOUR),
			OPEN24HOUR:      fmt.Sprintf("%s %.2f", model.TOSYMBOLDISPLAY, model.OPEN24HOUR),
			VOLUME24HOUR:    fmt.Sprintf("%s %.2f", model.FROMSYMBOLDISPLAY, model.VOLUME24HOUR),
			VOLUME24HOURTO:  fmt.Sprintf("%s %.2f", model.TOSYMBOLDISPLAY, model.VOLUME24HOURTO),
			LOW24HOUR:       fmt.Sprintf("%s %.2f", model.TOSYMBOLDISPLAY, model.LOW24HOUR),
			HIGH24HOUR:      fmt.Sprintf("%s %.2f", model.TOSYMBOLDISPLAY, model.HIGH24HOUR),
			PRICE:           fmt.Sprintf("%s %.2f", model.TOSYMBOLDISPLAY, model.PRICE),
			FROMSYMBOL:      model.FROMSYMBOLDISPLAY,
			TOSYMBOL:        model.TOSYMBOLDISPLAY,
			LASTUPDATE:      fmt.Sprint(model.LASTUPDATE),
			SUPPLY:          fmt.Sprintf("%s %d", model.FROMSYMBOLDISPLAY, model.SUPPLY),
			MKTCAP:          fmt.Sprintf("%s %.2f", model.TOSYMBOLDISPLAY, model.MKTCAP),
		}
	}
	return prices
}

func (s *cryptoServiceImpl) getLastPriceModels(prices types.PricesResponse) []types.LastPriceModel {
	models := make([]types.LastPriceModel, 0)
	for fsym, pair := range prices.RAW {
		for tsym, rawdata := range pair {
			model := types.LastPriceModel{
				FROMSYMBOL: fsym,
				TOSYMBOL: tsym,
				FROMSYMBOLDISPLAY: prices.DISPLAY[fsym][tsym].FROMSYMBOL,
				TOSYMBOLDISPLAY: prices.DISPLAY[fsym][tsym].TOSYMBOL,
				CHANGE24HOUR: rawdata.CHANGE24HOUR,
				CHANGEPCT24HOUR: rawdata.CHANGEPCT24HOUR,
				OPEN24HOUR: rawdata.OPEN24HOUR,
				VOLUME24HOUR: rawdata.VOLUME24HOUR,
				VOLUME24HOURTO: rawdata.VOLUME24HOURTO,
				LOW24HOUR: rawdata.LOW24HOUR,
				HIGH24HOUR: rawdata.HIGH24HOUR,
				PRICE: rawdata.PRICE,
				LASTUPDATE: rawdata.LASTUPDATE,
				SUPPLY: rawdata.SUPPLY,
				MKTCAP: rawdata.MKTCAP,
			}
			models = append(models, model)
		}
	}
	return models
}
