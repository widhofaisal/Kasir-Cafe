package config

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"kasir/cafe/model"
)

var(
	DB *gorm.DB
)

func init(){
	InitDB()
	InitialMigration()
}

type Config struct{
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB(){
	config := Config{
		DB_Username: "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "kasircafe_db",
	}

	// config for connect rds 
	// use this when deploy using docker
	// config_rds := Config{
	// 	DB_Username: "ownerwidho",
	// 	DB_Password: "kasircafeceria123",
	// 	DB_Port:     "3306",
	// 	DB_Host:     "db-kasircafe.ctbpmx2n7zth.ap-southeast-2.rds.amazonaws.com",
	// 	DB_Name:     "db_kasircafe",
	// }

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&model.Admin{}, &model.Product{}, &model.Cart{})
}