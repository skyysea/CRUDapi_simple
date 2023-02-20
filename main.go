package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int     `json: "id"`
	Isbn   string  `json: "isbn"`
	Title  string  `json: "title"`
	Author *Author `json: "author"`
}

type Author struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "firstname"`
}

var books []Book

// All book
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// single Book
func getsingleBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID, _ = strconv.Atoi(params["id"]); item.ID == item.ID {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create Book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)

	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

// Update Book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID, _ = strconv.Atoi(params["id"]); item.ID == item.ID {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = item.ID
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID, _ = strconv.Atoi(params["id"]); item.ID == item.ID {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	r := mux.NewRouter()

	books = append(books, Book{ID: 1, Isbn: "448743", Title: "Harry", Author: &Author{Firstname: "Jason", Lastname: "King"}})
	books = append(books, Book{ID: 2, Isbn: "448744", Title: "Potter", Author: &Author{Firstname: "JK.", Lastname: "Rolling"}})

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getsingleBook).Methods("GET")
	r.HandleFunc("/api/books", CreateBook).Methods("Post")
	r.HandleFunc("/api/books/{id}", UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
