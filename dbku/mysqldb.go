package dbku

import (
	"fmt"
	"tes1/model"
	"tes1/varglobal"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		varglobal.DatabaseUser,
		varglobal.DatabasePassword,
		varglobal.DatabaseHost,
		varglobal.DatabasePort,
		varglobal.DatabaseName,
	)

	var err error
	varglobal.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	fmt.Println("Database connected successfully")

	// Migrate the schema, if you have models to migrate
	// ini untuk migrasi skema table awal biar automatis
	varglobal.DB.AutoMigrate(&model.Book{})
	varglobal.DB.AutoMigrate(&model.InventoryItem{})
	fmt.Println("Database migration completed")

}
