package main

import (
	"diary/diary"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api", diary.RootHandler)

	log.Println("Server is runnig on port 8999!")
	http.ListenAndServe(":8999", nil)
}
