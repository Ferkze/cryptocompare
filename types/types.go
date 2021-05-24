package types

type SymbolPair struct {
	FROMSYMBOL string `json:"FROMSYMBOL"` // BTC
	TOSYMBOL   string `json:"TOSYMBOL"`   // USD
}

type LastPriceModel struct {
	FROMSYMBOL        string  `json:"FROMSYMBOL"`        // BTC
	TOSYMBOL          string  `json:"TOSYMBOL"`          // USD
	FROMSYMBOLDISPLAY string  `json:"FROMSYMBOLDISPLAY"` // Ƀ
	TOSYMBOLDISPLAY   string  `json:"TOSYMBOLDISPLAY"`   // $
	CHANGE24HOUR      float64 `json:"CHANGE24HOUR"`
	CHANGEPCT24HOUR   float64 `json:"CHANGEPCT24HOUR"`
	OPEN24HOUR        float64 `json:"OPEN24HOUR"`
	VOLUME24HOUR      float64 `json:"VOLUME24HOUR"`
	VOLUME24HOURTO    float64 `json:"VOLUME24HOURTO"`
	LOW24HOUR         float64 `json:"LOW24HOUR"`
	HIGH24HOUR        float64 `json:"HIGH24HOUR"`
	PRICE             float64 `json:"PRICE"`
	LASTUPDATE        int64   `json:"LASTUPDATE"`
	SUPPLY            int64   `json:"SUPPLY"`
	MKTCAP            float64 `json:"MKTCAP"`
}

type PriceRaw struct {
	CHANGE24HOUR    float64 `json:"CHANGE24HOUR"`
	CHANGEPCT24HOUR float64 `json:"CHANGEPCT24HOUR"`
	OPEN24HOUR      float64 `json:"OPEN24HOUR"`
	VOLUME24HOUR    float64 `json:"VOLUME24HOUR"`
	VOLUME24HOURTO  float64 `json:"VOLUME24HOURTO"`
	LOW24HOUR       float64 `json:"LOW24HOUR"`
	HIGH24HOUR      float64 `json:"HIGH24HOUR"`
	PRICE           float64 `json:"PRICE"`
	LASTUPDATE      int64   `json:"LASTUPDATE"`
	SUPPLY          int64   `json:"SUPPLY"`
	MKTCAP          float64 `json:"MKTCAP"`
}
type PriceDisplay struct {
	CHANGE24HOUR    string `json:"CHANGE24HOUR"`
	CHANGEPCT24HOUR string `json:"CHANGEPCT24HOUR"`
	OPEN24HOUR      string `json:"OPEN24HOUR"`
	VOLUME24HOUR    string `json:"VOLUME24HOUR"`
	VOLUME24HOURTO  string `json:"VOLUME24HOURTO"`
	LOW24HOUR       string `json:"LOW24HOUR"`
	HIGH24HOUR      string `json:"HIGH24HOUR"`
	PRICE           string `json:"PRICE"`

	FROMSYMBOL string `json:"FROMSYMBOL"` // BTC
	TOSYMBOL   string `json:"TOSYMBOL"`   // USD

	LASTUPDATE string `json:"LASTUPDATE"`
	SUPPLY     string `json:"SUPPLY"`
	MKTCAP     string `json:"MKTCAP"`
}

type PricesResponse struct {
	RAW     map[string]map[string]PriceRaw     `json:"RAW"`
	DISPLAY map[string]map[string]PriceDisplay `json:"DISPLAY"`
}
