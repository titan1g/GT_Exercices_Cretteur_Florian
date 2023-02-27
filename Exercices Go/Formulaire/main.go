package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			name := r.FormValue("name")
			email := r.FormValue("email")
			fmt.Printf("Name: %s\nEmail: %s\n", name, email)
		}
		fmt.Fprint(w, "<html><body><form method='post'><label for='name'>Name:</label><input type='text' name='name'><br><label for='email'>Email:</label><input type='email' name='email'><br><input type='submit' value='Submit'></form></body></html>")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
