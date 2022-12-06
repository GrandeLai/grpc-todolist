package dao

func migration() error {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&Task{})
	if err != nil {
		return err
	}
	return nil
}
