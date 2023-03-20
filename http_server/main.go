package main

import (
	"fmt"
	"strconv"
	"net/http"
)

const SERVER_PORT = 8080

// http://localhost:8090

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v", r)
	// fmt.Fprintf(w, "Welcome to my website!")
	htdocsFs := http.FileServer(http.Dir("htdocs/"))
	staticFs := http.FileServer(http.Dir("static/"))
	http.Handle("/public/", http.StripPrefix("/public/", htdocsFs))
	http.Handle("/static/", http.StripPrefix("/static/", staticFs))
}

func main() {
	port := strconv.Itoa(SERVER_PORT)
	fmt.Print("Port ", port, " Listening.")
	http.HandleFunc("/", requestHandler)

	http.ListenAndServe(":" + port, nil)
}
