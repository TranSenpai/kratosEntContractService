package data

import (
	"dormitory/internal/conf"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseConfig(data *conf.Data) *conf.Data_Database {
	return data.GetDatabase()
}

func NewMySQL(c *conf.Data_Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Dbname,
		c.Params,
	)
	// GORM can return specific errors related to the database dialect being used,
	// when TranslateError is enabled, GORM converts database-specific errors
	// into its own generalized errors.
	return gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true})
}
