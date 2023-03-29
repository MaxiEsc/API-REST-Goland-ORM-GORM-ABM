package models

import (
	"gorm/db"
)

type Usuario struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Usuarios []Usuario

func MigrarUsuario() {
	db.BaseDatos.AutoMigrate(Usuario{})
}