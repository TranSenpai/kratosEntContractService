package data

import (
	"dormitory/internal/ent"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewData, NewContractRepo, NewEntClient, NewDatabaseConfig)

type Data struct {
	dbPostgres *ent.Client
	log        *log.Helper
}

func (d *Data) DB() *ent.Client {
	return d.dbPostgres
}

func NewData(database *ent.Client, logger log.Logger) (*Data, func(), error) {
	logHelper := log.NewHelper(log.With(logger, "module", "contract-service/data"))
	// Clean-up function
	cleanup := func() {
		database.Close()
	}
	data := &Data{
		dbPostgres: database,
		log:        logHelper,
	}

	return data, cleanup, nil
}
