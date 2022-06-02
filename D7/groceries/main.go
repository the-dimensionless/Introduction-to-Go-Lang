package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/allgroceries", AllGroceries)                        // ----> To request all groceries
	r.HandleFunc("/groceries/{name}", SingleGrocery)                   // ----> To request a specific grocery
	r.HandleFunc("/groceries", GroceriesToBuy).Methods("POST")         // ----> To add  new grocery to buy
	r.HandleFunc("/groceries/{name}", UpdateGrocery).Methods("PUT")    // ----> To update a grocery
	r.HandleFunc("/groceries/{name}", DeleteGrocery).Methods("DELETE") // ----> Delete a grocery
	log.Fatal(http.ListenAndServe(":8080", r))

}
