package main

import (
	"fmt"
	"github-com/pteus/orders-api/middleware"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /hello", basicHandler)

	server := http.Server{
		Addr:    ":3000",
		Handler: middleware.Logging(router),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to start server", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
