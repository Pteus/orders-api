package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List")
}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetByID")
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateByID")
}

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteByID")
}
