package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/solve", solveHandler)
	http.Handle("/", http.FileServer(http.Dir("./public")))

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func solveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	value := r.FormValue("function")

	if value == "" {
		http.Error(w, "Please enter a value", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "form data: %s", value)
}
