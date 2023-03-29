package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func EnviarDatos(rw http.ResponseWriter, datos interface{}, estado int, mensaje string) {
	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(estado)

	salida, _ := json.Marshal(&datos)
	fmt.Fprintln(rw, string(salida))
}

func EnviarError(rw http.ResponseWriter, estado int, mensaje string) {
	rw.WriteHeader(estado)

	fmt.Fprintf(rw, "Fuente no encontrada")
}
