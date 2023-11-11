package dataservice

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"project1/model"
)

func CreateBook(db *sql.DB, w http.ResponseWriter, r *http.Request) error {

	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil { // newdecoder will create a instance. &book will store the decoded data. decoder json to go
		return err
	}

	query := `INSERT INTO books(title, author, year) VALUES (?,?,?)`
	_, err := db.Exec(query, book.Title, book.Author, book.Year)

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated) // to know if create suceesfully  status 201
	json.NewEncoder(w).Encode(book)   // book variable(go) to json format

	return nil

}
