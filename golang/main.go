package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Im Golang and im Running on port 9001")
	})

	http.ListenAndServe(":9001", nil)
	fmt.Println("Server running on 9001")
}
