package main

import (
	"log"
	// "fmt"
	"html/template"
	"net/http"
	"time"
	"encoding/json"
	"fmt"
	"html"

	// "github.com/genki28/app/cloud_functions"
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

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func HelloHTTP(w http.ResponseWriter, r *http.Request) {
        var d struct {
                Name string `json:"name"`
        }
        if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
                fmt.Fprint(w, "Hello, World!")
                return
        }
        if d.Name == "" {
                fmt.Fprint(w, "Hello, World!")
                return
        }
        fmt.Fprintf(w, "Hello, %s!", html.EscapeString(d.Name))
}

func main() {
	http.HandleFunc("/", testHandler)
	// http.HandleFunc("/cloudFunctions", cloud_functions.HelloHttp)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}