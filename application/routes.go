package application

import (
	"github-com/pteus/orders-api/handler"
	"net/http"
)

func loadRoutes() *http.ServeMux {
	router := http.NewServeMux()
	orderHandler := &handler.Order{}

	router.HandleFunc("GET /order", orderHandler.List)
	router.HandleFunc("POST /order", orderHandler.Create)
	router.HandleFunc("GET /order/{id}", orderHandler.GetByID)
	router.HandleFunc("PUT /order/{id}", orderHandler.UpdateByID)
	router.HandleFunc("DELETE /order/{id}", orderHandler.DeleteByID)

	return router
}
