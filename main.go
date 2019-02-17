package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/alecthomas/template"
	"github.com/michaeljoyner/breathe/air"
)

func handler(w http.ResponseWriter, r *http.Request) {
	report, err := air.GetReport()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprint(w, err)
		return
	}
	tpl, err := ioutil.ReadFile("views/base.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	t, err := template.New("webpage").Parse(string(tpl))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	t.Execute(w, report)
}

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3456", nil))
}
