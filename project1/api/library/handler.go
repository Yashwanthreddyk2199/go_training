// mostly used to check if we got the correct request method
package api

import (
	"database/sql"
	"net/http"
)

func CreateHandler(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) { // this style of func is anoynomus function
		if r.Method != http.MethodPost {
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
			return // it is only check for correct req method so we will not return anything
		}

		if err := CreateBookLogic(db, w, r); err != nil { // function createbooklogic in bizlogic
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}

}
