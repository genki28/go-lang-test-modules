package main

import (
	"log"
	// "fmt"
	"html/template"
	"net/http"
	"time"
)

func loadTemplate(name string) *template.Template {
	t, err := template.ParseFiles("templates/"+name+".html", "templates/_header.html", "templates/_footer.html")

	if err != nil {
		log.Fatalf("Not found file: %v", err)
	}

	return t
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")

	if err != nil {
		log.Fatalf("Not found file: %v", err)
	}

	if err := t.Execute(w, struct {
		Title string
		Message string
		Time time.Time
	}{
		Title: "テストしません。",
		Message: "Hello world",
		Time: time.Now(),
	}); err != nil {
		log.Printf("failed to execute template:", err)
	}
}

func main() {
	http.HandleFunc("/", testHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}