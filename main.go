package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/alecthomas/template"
	"github.com/michaeljoyner/breathe/air"

	_ "github.com/joho/godotenv/autoload"
)

func handler(w http.ResponseWriter, r *http.Request) {
	report, err := air.GetReport()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	tpl, err := ioutil.ReadFile("views/base.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	t, err := template.New("webpage").Parse(string(tpl))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	t.Execute(w, report)
}

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
