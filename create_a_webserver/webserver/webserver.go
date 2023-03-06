package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"

	"batsy.com/wiki"
)

const PUBLIC_DIR = "./public"

var templates = template.Must(template.ParseFiles(PUBLIC_DIR+"/template.html", PUBLIC_DIR+"/view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-z0-9]+)$")

func handlerFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, "This is so awesome %s", req.URL.Path[1:])
}

// remove duplicacy

func renderTemplate(htmlFileLocation string, writer http.ResponseWriter, page *wiki.Page) {
	err := templates.ExecuteTemplate(writer, htmlFileLocation, page)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func getTitle(writer http.ResponseWriter, req *http.Request) (string, error) {
	matched := validPath.FindStringSubmatch(req.URL.Path)

	if matched == nil {
		http.NotFound(writer, req)
		return "", errors.New("invalid page Title")
	}
	return matched[2], nil
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		matched := validPath.FindStringSubmatch(req.URL.Path)

		if matched == nil {
			http.NotFound(w, req)
			return
		}
		fn(w, req, matched[2])
	}
}

func viewHandler(w http.ResponseWriter, req *http.Request, title string) {
	p, err := wiki.LoadPage(title)

	if err != nil {
		http.Redirect(w, req, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate("view.html", w, p)
}

func editHandler(w http.ResponseWriter, req *http.Request, title string) {
	page, err := wiki.LoadPage(title)

	if err != nil {
		page = &wiki.Page{Title: "Tempfile"}
	}

	renderTemplate("template.html", w, page)
}

func saveHandler(w http.ResponseWriter, req *http.Request, title string) {
	body := req.FormValue("body")

	page := &wiki.Page{Title: title, Body: []byte(body)}
	err := page.Save()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, req, "/view/"+title, http.StatusFound)
}

func main() {

	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
