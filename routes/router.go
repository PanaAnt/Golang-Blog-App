package routes

import (
	"blogApp/handlers"
	"blogApp/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handlers.UserHandler, postHandler *handlers.PostHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/register", userHandler.Register).Methods("POST")

	// sub router for protected routes
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// auth routes
	protected.HandleFunc("/protect", handlers.ProtectedHandler).Methods("GET")
	protected.HandleFunc("/protect", handlers.ProtectedHandler).Methods("GET")
	protected.HandleFunc("/posts/{id}", postHandler.GetPostByID).Methods("GET")
	protected.HandleFunc("/posts", postHandler.GetAllPosts).Methods("GET")
	protected.HandleFunc("/posts", postHandler.CreatePost).Methods("POST")
	protected.HandleFunc("/posts/:id", postHandler.UpdatePost).Methods("PUT")
	protected.HandleFunc("/posts/:id", postHandler.DeletePost).Methods("DELETE")

	return r
}
