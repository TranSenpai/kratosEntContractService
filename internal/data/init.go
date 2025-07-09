package data

import (
	"context"
)

func (data *Data) InitSchema() error {
	if err := createTableContract(data.DB(), context.Background()); err != nil {
		return err
	}

	data.log.Info("Schema initialized")
	return nil
}
