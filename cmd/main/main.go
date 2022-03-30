package main

import (
	"fmt"
	"go-bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	r := mux.NewRouter()
	routes.BookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Server started at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
