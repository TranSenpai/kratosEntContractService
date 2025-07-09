package data

import (
	"context"
	"dormitory/internal/ent"
	"fmt"
	_ "github.com/lib/pq"
)

func createTableContract(client *ent.Client, ctx context.Context) error {
	if err := client.Schema.Create(ctx); err != nil {
		return fmt.Errorf("failed creating schema resources: %w", err)
	}

	return nil
}
