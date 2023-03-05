package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"batsy.com/wiki"
)

func handlerFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, "This is so awesome %s", req.URL.Path[1:])
}

// remove duplicacy

func renderTemplate(htmlFileLocation string, writer http.ResponseWriter, page *wiki.Page) {
	respTemplate, err := template.ParseFiles(htmlFileLocation)

	if err != nil {
		log.Fatal("Something went wrong", err)
	}

	respTemplate.Execute(writer, page)
}

func viewHandler(w http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/view/"):]

	p, err := wiki.LoadPage(title)

	if err != nil {
		http.Redirect(w, req, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate("./public/view.html", w, p)
}

func editHandler(w http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/edit/"):]

	page, err := wiki.LoadPage(title)

	if err != nil {
		page = &wiki.Page{Title: "Tempfile"}
	}

	renderTemplate("./public/template.html", w, page)
}

func saveHandler(w http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/save/"):]
	body := req.FormValue("body")

	page := &wiki.Page{Title: title, Body: []byte(body)}
	page.Save()

	http.Redirect(w, req, "/view/"+title, http.StatusFound)
}

func main() {

	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
