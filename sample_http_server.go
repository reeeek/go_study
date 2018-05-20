package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Hello, World!!")
		w.Write([]byte("Hellor, World!!!"))
	})
	http.ListenAndServe(":8080", nil)
}
