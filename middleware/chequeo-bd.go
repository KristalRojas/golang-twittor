package middleware

import (
	"net/http"

	"github.com/KristalRojas/golang-twittor/database"
)

//ChequeoBD public
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if database.ChequeoConexion() == 0 {
			http.Error(rw, "Conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
