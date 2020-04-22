package middlew

import (
	"PaginaWebGoReact/pkg/routers"
	"net/http"
)

func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		_, _, _, err := routers.ProcesoToken(req.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el token ! "+ err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, req)
	}
}