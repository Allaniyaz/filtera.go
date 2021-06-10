package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func myHandleFunc() {
	//http.Handle("/static", http.FileServer(http.Dir("static")))
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	html.ExecuteTemplate(w, "index", nil)
}

func main() {
	myHandleFunc()
}
