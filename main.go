package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func echo(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status  int       `json:"status`
		Time    time.Time `json:"time`
		Message string    `json:"message`
	}

	w.Header().Set("Content-Type", "application/json")

	res := &Response{
		Status:  http.StatusOK,
		Time:    time.Now(),
		Message: "Json API Server by Golang on Docker",
	}

	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/", echo)
	http.ListenAndServe(":8888", nil)
}
