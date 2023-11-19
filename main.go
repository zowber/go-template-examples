package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title   string
	Content string
	Meta    map[string]string
}

var data PageData = PageData{
	Title:   "Cool page",
	Content: "Cool content",
	Meta: map[string]string{
		"Cool key1": "Cool val1",
		"Cool key2": "Cool val2",
	},
}

func main() {
	// stdlib html/template
	go func() {
		var indexHandler = func(w http.ResponseWriter, r *http.Request) {
			tmpl := template.Must(template.ParseGlob("./templates/*"))
			tmpl.Execute(w, data)
		}

		mux := http.NewServeMux()
		mux.HandleFunc("/", indexHandler)
		log.Println("Coming up on port 3000")
		err := http.ListenAndServe(":3000", mux)
		if err != nil {
			log.Fatal("Oops! The http server crashed ;(", err)
		}
	}()

	// templ
	go func() {
		var indexHandler = func(w http.ResponseWriter, r *http.Request) {
			page := index(data)
			page.Render(context.Background(), w)
		}

		mux := http.NewServeMux()
		mux.HandleFunc("/", indexHandler)
		log.Println("Coming up on port 3001")
		err := http.ListenAndServe(":3001", mux)
		if err != nil {
			log.Fatal("Oops! The http server crashed ;(", err)
		}
	}()

	select {}
}
