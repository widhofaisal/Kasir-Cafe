package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kasir/cafe/model"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	// <start> FOR DEVELOPMENT >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// config := Config{
	// 	DB_Username: "root",
	// 	DB_Password: "",
	// 	DB_Port:     "3306",
	// 	DB_Host:     "localhost",
	// 	DB_Name:     "kasircafe_db",
	// }

	// connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	config.DB_Username,
	// 	config.DB_Password,
	// 	config.DB_Host,
	// 	config.DB_Port,
	// 	config.DB_Name,
	// )
	// <end> FOR DEVELOPMENT >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	// <start> FOR PRODUCTION >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	config_rds := Config{
		DB_Username: "ownerwidho",
		DB_Password: "kasircafeceria123",
		DB_Port:     "3306",
		DB_Host:     "db-kasircafe.ctbpmx2n7zth.ap-southeast-2.rds.amazonaws.com",
		DB_Name:     "db_kasircafe",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config_rds.DB_Username,
		config_rds.DB_Password,
		config_rds.DB_Host,
		config_rds.DB_Port,
		config_rds.DB_Name,
	)
	// <end> FOR PRODUCTION >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&model.Admin{}, &model.Product{}, &model.Cart{}, &model.Payment{})
}
