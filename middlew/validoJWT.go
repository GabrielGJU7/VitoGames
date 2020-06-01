package middlew

import (
	"net/http"

	"github.com/GabrielGJU7/VitoGames/routers"
)

/*ValidoJWT comentario*/
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el token ! "+err.Error(), http.StatusAccepted)
			return
		}
		next.ServeHTTP(w, r)
	}
}
