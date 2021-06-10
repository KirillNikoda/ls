package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KirillNikoda/lowskill/models"
)

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func GetAllBooks(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Get infos about all books in database")
	writer.WriteHeader(200)
	log.Println(models.DB)
	json.NewEncoder(writer).Encode(models.DB)
}
