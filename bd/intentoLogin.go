package bd

import (
	"github.com/GabrielGJU7/VitoGames/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin sirve para validar el inicio de sesion*/
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}

	paswordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, paswordBytes)

	if err != nil {
		return usu, false
	}

	return usu, true
}
