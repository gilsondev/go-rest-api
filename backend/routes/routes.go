package routes

import (
	"log"
	"net/http"

	"github.com/gilsondev/go-rest-api/controllers"
	"github.com/gilsondev/go-rest-api/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

  r.Use(middleware.DefineContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.AtualizaPersonalidade).Methods("Put")
	r.HandleFunc("/api/personalidades/{id}", controllers.RemovePersonalidade).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
