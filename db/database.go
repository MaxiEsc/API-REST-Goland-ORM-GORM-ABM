package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Paquete ORM para Go
var dsn = "root:12345678@tcp(localhost:3306)/goweb_db?charset=utf8mb4&parseTime=True&loc=Local"

//Funcion anonima que se autoinicializa
var BaseDatos = func() (db *gorm.DB) {
	if db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{}); error != nil {
		fmt.Println("Error en la conexión", error)
		panic(error)
	} else {
		fmt.Println("Conexión Exitosa")
		return db
	}
}()