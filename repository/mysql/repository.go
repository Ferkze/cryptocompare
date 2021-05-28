package mysql

import (
	"fmt"

	"github.com/ferkze/cryptocompare/repository"
	"github.com/ferkze/cryptocompare/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LastPriceModel struct {
	FROMSYMBOL        string  `json:"FROMSYMBOL" gorm:"primaryKey; not null"` // BTC
	TOSYMBOL          string  `json:"TOSYMBOL" gorm:"primaryKey; not null"`   // USD
	FROMSYMBOLDISPLAY string  `json:"FROMSYMBOLDISPLAY"`                      // Ƀ
	TOSYMBOLDISPLAY   string  `json:"TOSYMBOLDISPLAY"`                        // $
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

type mysqlCrytoRepository struct {
	db *gorm.DB
}

func NewCryptoRepository(dsn string) repository.CryptoRepository {
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(LastPriceModel{})
	return &mysqlCrytoRepository{db}
}

func (r *mysqlCrytoRepository) GetLastPrice(fsyms, tsyms []string) ([]types.LastPriceModel, error) {
	var models []types.LastPriceModel
	result := r.db.Where("fromsymbol IN ?", fsyms).
		Where("tosymbol IN ?", fsyms).
		Order("lastupdate desc").
		Find(&models)
	return models, result.Error
}
func (r *mysqlCrytoRepository) GetSymbols() ([]types.LastPriceModel, error) {
	var models []types.LastPriceModel
	result := r.db.Select("FROMSYMBOL", "TOSYMBOL").Find(&models)
	return models, result.Error
}
func (r *mysqlCrytoRepository) BulkUpdateLastPrices(models []types.LastPriceModel) error {
	result := r.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&models)
	return result.Error
}
