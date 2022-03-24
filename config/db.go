package config

import (
	"ProjectSavePassword/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	DNS := "user=postgres password=adityarizky1020 host=db.jgjyjvyldoamqndazixl.supabase.co TimeZone=Asia/Singapore port=5432 dbname=postgres "
	db, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(model.SaveData{}, model.DataUser{})
	return db, err
}
