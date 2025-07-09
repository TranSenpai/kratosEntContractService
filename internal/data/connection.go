package data

import (
	"dormitory/internal/conf"
	ent "dormitory/internal/ent"
	"fmt"
)

func NewDatabaseConfig(data *conf.Data) *conf.Data_Database {
	return data.GetDatabase()
}

func NewEntClient(c *conf.Data_Database) (*ent.Client, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Dbname,
		c.Params,
	)

	return ent.Open("postgres", dsn)
}
