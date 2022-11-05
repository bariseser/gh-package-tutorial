package main

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func main() {
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		object := response{Status: "success", Code: 200}
		resp, _ := json.Marshal(object)
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(resp)
	})
	http.ListenAndServe(":9999", nil)
}
