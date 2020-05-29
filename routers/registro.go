package routers

import (
	"encoding/json"
	"net/http"

	"github.com/GabrielGJU7/VitoGames/bd"
	"github.com/GabrielGJU7/VitoGames/models"
)

/*Registro es la funcion para crear en la BD el registro de usuarios*/

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El Email de usuario es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "El password debe tener mas de 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario con ese Email registrado", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al rgistrar los datos"+err.Error(), 400)
	}

	if status == false {
		http.Error(w, "Ocurrio un error en el regitro"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
