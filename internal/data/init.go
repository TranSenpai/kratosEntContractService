package data

func (d *Data) InitSchema() error {
	if err := createTableContract(d.db); err != nil {
		return err
	}
	createPartitionContract(d.db)

	d.log.Info("Schema initialized")
	return nil
}
