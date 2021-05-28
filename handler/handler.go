package handler

import (
	"net/http"
	"strings"

	"github.com/ferkze/cryptocompare/handler/utils"
	"github.com/ferkze/cryptocompare/repository"
	"github.com/ferkze/cryptocompare/services"
)

type CryptoHandler struct {
	repo repository.CryptoRepository
	service services.CryptoService
}

func NewCryptoHandler(repo repository.CryptoRepository, service services.CryptoService) *CryptoHandler {
	return &CryptoHandler{repo, service}
}

func (h *CryptoHandler) GetPairPrices(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	qfsyms := strings.TrimSpace(queries.Get("fsyms"))
	qtsyms := strings.TrimSpace(queries.Get("tsyms"))
	if qfsyms == "" || qtsyms == "" {
		utils.JSON(w, http.StatusBadRequest, map[string]string{"message": "No given pairs"})
		return
	}
	fsyms := strings.Split(qfsyms, ",")
	tsyms := strings.Split(qtsyms, ",")
	response, err := h.service.GetLastSymbolsPrice(fsyms, tsyms)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, map[string]string{"message": "Internal err"})
		return
	}
	utils.JSON(w, http.StatusOK, response)
}