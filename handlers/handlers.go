package handlers

import (
	"encoding/json"
	"fmt"
	"gorm/db"
	"gorm/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

const USER_ELIMINADO = "Se ha Eliminado el Usuario correctamente"
const USER_AGREGADO = "Se ha Creado el Usuario correctamente"
const USER_MODIFICADO = "Se ha Modificado el Usuario correctamente"
const USER_LISTADO = "Se han listado Los Usuarios correctamente"
const USER_BUSCADO = "Se ha encontrado el Usuario correctamente"
const USER_ERROR = "Error en la comunicacion con la base de datos"

func GetUsuarios(rw http.ResponseWriter, r *http.Request) {

	usuarios := models.Usuarios{}
	db.BaseDatos.Find(&usuarios)
	EnviarDatos(rw, usuarios, http.StatusOK, USER_LISTADO)
}

func GetUsuario(rw http.ResponseWriter, r *http.Request) {
	if usuario, error := obtenerUsuarioPorId(r); error != nil {
		EnviarError(rw, http.StatusNotFound, USER_ERROR)
	} else {
		EnviarDatos(rw, usuario, http.StatusOK, USER_BUSCADO)
	}
}

func CreateUsuario(rw http.ResponseWriter, r *http.Request) {

	//Obtener el registro
	usuario := models.Usuario{}
	decodificar := json.NewDecoder(r.Body)

	if error := decodificar.Decode(&usuario); error != nil {
		EnviarError(rw, http.StatusUnprocessableEntity, USER_ERROR)
		fmt.Fprintln(rw, "Error")
	} else {
		db.BaseDatos.Save(&usuario)
		EnviarDatos(rw, usuario, http.StatusCreated, USER_AGREGADO)
		fmt.Fprintln(rw, "Proceso realizado correctamente")
	}
}

func UpdateUsuario(rw http.ResponseWriter, r *http.Request) {

	//Obtener registro
	var usId int64

	if usuario_antiguo, error := obtenerUsuarioPorId(r); error != nil {
		EnviarError(rw, http.StatusNotFound, USER_ERROR)
	} else {
		usId = usuario_antiguo.Id

		usuario := models.Usuario{}
		decodificar := json.NewDecoder(r.Body)

		if error := decodificar.Decode(&usuario); error != nil {
			EnviarError(rw, http.StatusUnprocessableEntity, USER_ERROR)
			fmt.Fprintln(rw, "Error")
		} else {
			usuario.Id = usId
			db.BaseDatos.Save(&usuario)
			EnviarDatos(rw, usuario, http.StatusAccepted, USER_MODIFICADO)
			fmt.Fprintln(rw, "Proceso realizado correctamente")
		}
	}
}

func DeleteUsuario(rw http.ResponseWriter, r *http.Request) {
	if usuario, error := obtenerUsuarioPorId(r); error != nil {
		EnviarError(rw, http.StatusNotFound, USER_ELIMINADO)
	} else {
		db.BaseDatos.Delete(&usuario)
		EnviarDatos(rw, usuario, http.StatusOK, USER_ELIMINADO)
	}
}

func obtenerUsuarioPorId(r *http.Request) (models.Usuario, *gorm.DB) {
	//obtener el id
	v := mux.Vars(r)
	usuarioId, _ := strconv.Atoi(v["id"])

	usuario := models.Usuario{}
	if error := db.BaseDatos.First(&usuario, usuarioId); error.Error != nil {
		return usuario, error
	} else {
		return usuario, nil
	}
}
