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
	id, err := strconv.Atoi(mux.Vars(request)["id"])

	if err != nil {
		msg := models.Message{Message: "error"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	var book models.Book
	var found bool

	for _, b := range models.DB {
		if b.ID == id {
			book = b
			found = true
		}
	}

	if !found {
		msg := models.Message{Message: "No such book in database"}
		json.NewEncoder(writer).Encode(msg)
	}

	json.NewEncoder(writer).Encode(book)
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	var newBook models.Book
	err := json.NewDecoder(request.Body).Decode(&newBook)
	if err != nil {
		writer.WriteHeader(400)
		msg := models.Message{Message: "something went wrong"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	models.DB = append(models.DB, newBook)
	json.NewEncoder(writer).Encode(newBook)
	log.Println("Successfully created")
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
