package diary

import (
	"diary/diary/pgdb"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
	}

	var task pgdb.Tasks

	json.Unmarshal(data, &task)

	switch r.Method {
	case http.MethodGet:
	case http.MethodPost:
	case http.MethodDelete:
	case http.MethodPut:
	}
}

func DiaryHandler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
	}

	var diary pgdb.Diary

	json.Unmarshal(data, &diary)

	switch r.Method {
	case http.MethodGet:
	case http.MethodPost:
	case http.MethodDelete:
	case http.MethodPut:
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request) {

}

func AuthHandler(w http.ResponseWriter, r *http.Request) {

}
