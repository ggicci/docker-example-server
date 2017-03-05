package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func index(rw http.ResponseWriter, r *http.Request) {
	io.WriteString(rw, "Welcome to an example server in docker!")
}

func health(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(map[string]interface{}{
		"time":    time.Now(),
		"message": "ok",
	})
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/v1/health", health)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
