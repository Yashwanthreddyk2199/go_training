package api

import (
	"database/sql"
	"net/http"
	"project1/dataservice"
)

func CreateBookLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	return dataservice.CreateBook(db, w, r)

}
