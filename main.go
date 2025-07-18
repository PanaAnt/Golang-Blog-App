package main

import (
	"blogApp/db"
	"blogApp/handlers"
	"blogApp/repository"
	"blogApp/routes"
	"blogApp/services"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db.InitDb()

	//init repo
	userRepo := &repository.UserRepo{}
	postRepo := &repository.PostRepo{}

	//init service
	userService := &services.UserService{Repo: userRepo}
	postService := &services.PostService{Repo: postRepo}
	//init handlers
	userHandler := &handlers.UserHandler{Service: userService}
	postHandler := &handlers.PostHandler{Service: postService}
	
	//routes
	r := routes.SetupRouter(userHandler, postHandler)
	//start server

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Failed to start server", err)
	}
	fmt.Println("server started on :8080")

}
