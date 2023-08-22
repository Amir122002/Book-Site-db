package handlers

import (
	"book/internal/database"
	"book/pkg/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

var db, _ = database.DataBase()

func Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var requestBody map[string]interface{}
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	title, ok := requestBody["title"].(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	descriptions, ok := requestBody["descriptions"].(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	author, ok := requestBody["author"].(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db.Exec("insert into Book(title,descriptions,author) values($1,$2,$3)", title, descriptions, author)

}

func ReadId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var book []models.Book
	db.Raw("select * from Book where id=$1", id).Scan(&book)
	jsonBytes, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func Read(w http.ResponseWriter, r *http.Request) {
	var book []models.Book
	vars := mux.Vars(r)
	pageStr := vars["page"]
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	perPage := 10
	offset := (page - 1) * perPage
	db.Table("book").Limit(perPage).Offset(offset).Find(&book)
	//db.Limit(perPage).Offset(offset).Find(&book)
	jsonBytes, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var requestBody map[string]interface{}
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	title, ok := requestBody["title"].(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	descriptions, ok := requestBody["descriptions"].(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	author, ok := requestBody["author"].(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db.Exec("UPDATE Book set title=$1,descriptions=$2,author=$3,update_at=current_timestamp where id=$4 ", title, descriptions, author, id)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	db.Exec("delete from Book where id=$1", id)
	w.WriteHeader(http.StatusOK)
}
