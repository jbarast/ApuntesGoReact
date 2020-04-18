package handlers
import(
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func Manejadores(){
	router := mux.NewRouter()

	// Miro si hay una variable de entorno PORT.
	PORT := os.Getenv("PORT")
	if PORT == ""{
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router) // Para temas de permisos. Libreria cors.
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}