package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/damocles217/server/database"
	"github.com/damocles217/server/middlewares"
	"github.com/damocles217/server/router/user"
)

type App struct {
	Router     *mux.Router
	Collection *mongo.Collection
}

func (a *App) NewServer(uri string, port string) {
	// Creting database
	collection := database.MongoStart(uri)
	a.Collection = collection

	// Creating the router
	router := mux.NewRouter()

	// CORS
	router.Use(mux.CORSMethodMiddleware(router))

	// Middlewares

	router.Use(middlewares.AuthMiddleware())

	// Routes

	router.HandleFunc("/user", a.createUser).Methods(http.MethodPost)
	router.HandleFunc("/user/{id}", a.getUser).Methods(http.MethodGet)

	// Server on
	http.ListenAndServe(port, router)

	// end of the router
	a.Router = router

}

func (a *App) createUser(w http.ResponseWriter,
	r *http.Request) {
	user.MakeUser(w, r, a.Collection)
}

func (a *App) getUser(w http.ResponseWriter,
	r *http.Request) {
}
