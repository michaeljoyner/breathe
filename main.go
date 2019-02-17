package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/michaeljoyner/breathe/air"
)

func handler(w http.ResponseWriter, r *http.Request) {
	report, err := air.GetReport()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprint(w, err)
		return
	}

	// tpl, err := ioutil.ReadFile("views/base.html")
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	fmt.Fprint(w, err)
	// 	return
	// }
	// t, err := template.New("webpage").Parse(string(tpl))
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	fmt.Fprint(w, err)
	// 	return
	// }

	// err = t.Execute(w, report)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	fmt.Fprint(w, err)
	// }
	fmt.Fprintln(w, report.Warning)
}

func main() {
	godotenv.Load("/home/forge/air.magjoyner.com/.env")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3456", nil))
}
