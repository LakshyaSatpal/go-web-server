package main

import (
	"log"
	"net/http"
)

func main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, r.URL.Path[1:])
	//  this will show index.html at localhost:8080/static/index.html
	// })

	http.Handle("/", http.FileServer(http.Dir("./static"))) // to show index.html at localhost:8080/index.html

	log.Fatal(http.ListenAndServe(":8080", nil))
}
