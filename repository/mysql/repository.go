package mysql

import (
	"github.com/ferkze/cryptocompare/repository"
	"github.com/ferkze/cryptocompare/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
type mysqlCrytoRepository struct {
	db *gorm.DB
}

func NewCryptoRepository(dsn string) repository.CryptoRepository {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
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
func (r *mysqlCrytoRepository) GetSymbols() ([]types.SymbolPair, error) {
	return nil, nil
}
func (r *mysqlCrytoRepository) BulkUpdateLastPrices(models []types.LastPriceModel) error {
	return nil
}
