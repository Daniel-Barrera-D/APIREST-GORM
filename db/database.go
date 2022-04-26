package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Realiza conexión a la base de datos
var dsn = "root:28Z28E10U23S.@tcp(localhost:3306)/goweb_db?charset=utf8mb4&parseTime=True&loc=Local"
var Database = func() (db *gorm.DB) {
	if db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{}); error != nil {
		fmt.Println("Error en la conexión", error)
		panic(error)
	} else {
		fmt.Println("Conexión exitosa")
		return db
	}
}()