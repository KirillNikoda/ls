package main

import (
	"github.com/KirillNikoda/lowskill/utils"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	apiPrefix = "/api/v1"
)

var (
	port                    string
	bookResourcePrefix      = apiPrefix + "/book"  // api/v1/book
	manyBooksResourcePrefix = apiPrefix + "/books" // api/v1/books
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("could not find .env file:", err)
	}
	port = os.Getenv("app_port")
}

func main() {
	router := mux.NewRouter()

	utils.BuildBookResource(router, bookResourcePrefix)

	utils.BuildManyBooksResource(router, manyBooksResourcePrefix)

	log.Println("Router configured successfully! Let's go!")

	log.Fatal(http.ListenAndServe(":8080", router))
}
