package main

import (
	"fmt"
	"net/http"
)

const PORT int = 8090

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8090", nil)
	log.print("Port", PORT, "Listening")
}
