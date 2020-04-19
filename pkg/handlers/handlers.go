package handlers
import(
	"PaginaWebGoReact/pkg/middlew"
	"PaginaWebGoReact/pkg/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func Manejadores(){
	router := mux.NewRouter()

	// Router.
	router.HandleFunc("/registro",middlew.ChequeoDB(routers.Registro)).Methods("POST")
	router.HandleFunc("/login",middlew.ChequeoDB(routers.Login)).Methods("POST")

	// Miro si hay una variable de entorno PORT.
	PORT := os.Getenv("PORT")
	if PORT == ""{
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router) // Para temas de permisos. Libreria cors.
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}