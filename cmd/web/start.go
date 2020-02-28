package web

import (
	"github.com/gorilla/mux"
	"github.com/hiram66/user-service/cmd/web/controllers"
	_ "github.com/hiram66/user-service/storage"
	"log"
	"net/http"
)

func Start(port string) {
	r := mux.NewRouter()
	registerRoutes(r)
	log.Printf("listening port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func registerRoutes(r *mux.Router) {
	r.HandleFunc("/users", controllers.AddUser).Methods(http.MethodPost)
	r.HandleFunc("/users/{userId}", controllers.GetUserById).Methods(http.MethodGet)
	r.HandleFunc("/users", controllers.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/users/{userId}", controllers.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/users/{userId}", controllers.DeleteUser).Methods(http.MethodDelete)
}
