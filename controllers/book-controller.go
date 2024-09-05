package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/damaisme/go-mysql-bookstore-api/config"
	"github.com/damaisme/go-mysql-bookstore-api/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = config.SetupDatabaseConn()
	db.AutoMigrate(&models.Book{})
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	db.Find(&books)
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	db.Create(&book)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["ID"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	var book models.Book
	db.Where("ID=?", ID).Delete(&book)

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Book
	bookId := params["ID"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	db.Where("ID=?", ID).Find(&book)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["ID"]

	var updateBook models.Book
	_ = json.NewDecoder(r.Body).Decode(&updateBook)

	var bookDetails models.Book

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	result := db.Where("ID=?", ID).Find(&bookDetails)

	if result.Error != nil {
		fmt.Println("Error retrieving book:", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		fmt.Println("No book found")
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusNotModified)
		res, _ := json.Marshal(&bookDetails)
		w.Write(res)
		return
	}

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
