package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewData, NewContractRepo, NewMySQL, NewDatabaseConfig)

type Data struct {
	db  *gorm.DB
	log *log.Helper
}

func (d *Data) DB() *gorm.DB {
	return d.db
}

func NewData(database *gorm.DB, logger log.Logger) (*Data, func(), error) {
	logHelper := log.NewHelper(log.With(logger, "module", "contract-service/data"))
	// Clean-up function
	cleanup := func() {
		sqlDB, err := database.DB()
		if err == nil {
			sqlDB.Close()
		} else {
			logHelper.Error(err)
		}
	}
	data := &Data{
		db:  database,
		log: logHelper,
	}

	return data, cleanup, nil
}
