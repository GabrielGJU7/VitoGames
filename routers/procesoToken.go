package routers

import (
	"errors"
	"strings"

	"github.com/GabrielGJU7/VitoGames/bd"
	"github.com/GabrielGJU7/VitoGames/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email comentario*/
var Email string

/*IDUsuario comentario*/
var IDUsuario string

/*ProcesoToken proceso token pra extraer susvalores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MasterDelDesarrollo")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if tkn.Valid {
		return claims, false, string(""), errors.New("Token Invalido")
	}

	return claims, false, string(""), err
}
