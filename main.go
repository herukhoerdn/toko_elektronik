package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/toko-elektronik/config"
	"github.com/toko-elektronik/handler"
	"github.com/toko-elektronik/repository"
	"github.com/toko-elektronik/usecase"
)

func main() {
	db := config.InitDB()
	defer db.Close()

	repo := repository.NewRepository(db)
	uc := usecase.NewUsecase(repo)
	router := mux.NewRouter()
	handler.InitRoute(router, uc)

	// Add CORS support using gorilla/handlers.CORS
	headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "FETCH"})

	// Start the HTTP server
	fmt.Println("server is running on port :8080")
	log.Fatal(http.ListenAndServe("localhost:8080", handlers.CORS(headersOk, originsOk, methodsOk)(router)))

}
