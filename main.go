package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

const port string = ":8080"

func main() {

	fmt.Println("Initialzing Router ..")

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./src/template/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// fmt.Println("Hello World !!")
		err = tmpl.Execute(w, "Kiran")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})
	fmt.Printf("Listening at http://localhost:%s ", port[1:])
	http.ListenAndServe(port, router)
}
