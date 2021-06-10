package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/KirillNikoda/lowskill/models"
	"github.com/gorilla/mux"
)

func GetBookById(writer http.ResponseWriter, request *http.Request) {

}

func CreateBook(writer http.ResponseWriter, request *http.Request) {

}

func UpdateBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Updating book...")

	id, err := strconv.Atoi(mux.Vars(request)["id"])

	if err != nil {
		log.Println("error while parsing happened", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "don't use parameter ID as uncasted int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	var newBook models.Book

	oldBook, ok := models.FindBookById(id)

	if !ok {
		writer.WriteHeader(400)
		msg := models.Message{Message: "book with that ID does not exists in database"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	err = json.NewDecoder(request.Body).Decode(&newBook)

	if err != nil {
		writer.WriteHeader(400)
		msg := models.Message{Message: "provided json file is invalid"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	*oldBook = newBook
}

func DeleteBookById(writer http.ResponseWriter, request *http.Request) {

}
