package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/victorcel/chucknorris/handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		panic(err.Error())
	}
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/jokes", handlers.GetJacksHandle()).Methods(http.MethodGet)

	log.Println("Iniciando servidor en puerto:", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatal("Error: ", err)
	}
}
