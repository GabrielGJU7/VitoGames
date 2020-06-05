package routers

import (
	"encoding/json"
	"net/http"

	"github.com/GabrielGJU7/VitoGames/bd"
	"github.com/GabrielGJU7/VitoGames/models"
)

/*Modificar perfil comentario*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos Incorrectods"+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificr el registro, reitente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado modificar el registro del usuario"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
