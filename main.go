package main

import (
	"diary/diary"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api", diary.RootHandler)
	http.HandleFunc("/api/tasks", diary.TasksHandler)
	http.HandleFunc("/api/diary", diary.DiaryHandler)
	http.HandleFunc("/api/user", diary.UserHandler)
	http.HandleFunc("/api/auth", diary.AuthHandler)

	log.Println("Server is runnig on port 8992!")
	if err := http.ListenAndServe(":8992", nil); err != nil {
		log.Fatal(err)
	}
}
