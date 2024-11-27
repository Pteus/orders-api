package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", basicHandler)

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println("failed to start server", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
