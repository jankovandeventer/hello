package main

import "net/http"

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("jankovandeventer"))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<a href=\"mailto:janko@batao.io\">contact</a>"))
}
