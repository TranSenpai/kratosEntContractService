package data

import (
	entity "dormitory/internal/entities"

	"gorm.io/gorm"
)

func migrateTable(connection *gorm.DB) error {
	return connection.AutoMigrate(&entity.Contract{})
}

func createTableContract(dbConnection *gorm.DB) error {
	err := migrateTable(dbConnection)
	if err != nil {
		return err
	}

	return nil
}

func createPartitionContract(dbConnection *gorm.DB) {
	dbConnection.Debug().Model(&entity.Contract{}).Exec(`
	ALTER TABLE contracts
	ADD COLUMN registry_partition int unsigned GENERATED ALWAYS AS (cast(DATE(registry_at) / 100 as unsigned)) STORED`)

	dbConnection.Debug().Model(&entity.Contract{}).Exec(`
	ALTER TABLE contracts DROP PRIMARY KEY, ADD PRIMARY KEY (id, registry_partition)`)

	dbConnection.Debug().Model(&entity.Contract{}).
		Exec(`ALTER TABLE contracts PARTITION BY RANGE COLUMNS(registry_partition)(
			PARTITION p01 VALUES LESS THAN (202502),
			PARTITION p02 VALUES LESS THAN (202503),
			PARTITION p03 VALUES LESS THAN (202504),
			PARTITION p04 VALUES LESS THAN (202505),
			PARTITION p05 VALUES LESS THAN (202506),
			PARTITION p06 VALUES LESS THAN (202507),
			PARTITION p07 VALUES LESS THAN (202508),
			PARTITION p08 VALUES LESS THAN (202509),
			PARTITION p09 VALUES LESS THAN (202510),
			PARTITION p10 VALUES LESS THAN (202511),
			PARTITION p11 VALUES LESS THAN (202512),
			PARTITION p12 VALUES LESS THAN (202513))`)
}
