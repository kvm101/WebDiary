package diary

import (
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
	case http.MethodPost:
	case http.MethodDelete:
	case http.MethodPut:
	}
}

func DiaryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
	case http.MethodPost:
	case http.MethodDelete:
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request) {

}

func AuthHandler(w http.ResponseWriter, r *http.Request) {

}
