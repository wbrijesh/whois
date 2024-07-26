// package main
//
// import (
// 	"fmt"
//
// 	"github.com/likexian/whois"
// )
//
// func main() {
// 	result, _ := whois.Whois("brijesh.dev")
//
// 	fmt.Println("Result: ", result)
// }

package main

import (
	"html/template"
	"net/http"

	"github.com/likexian/whois"
)

type WHOISRecord struct {
	Domain string `json:"domain"`
	Result string `json:"result"`
}

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/lookup", whoisLookupHandler)
	http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func whoisLookupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else if r.Method != http.MethodPost {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}

	domain := r.FormValue("domain")
	result, err := whois.Whois(domain)
	if err != nil {
		http.Error(w, "Error fetching WHOIS data", http.StatusInternalServerError)
		return
	}

	whoisRecord := &WHOISRecord{
		Domain: domain,
		Result: result,
	}

	templates.ExecuteTemplate(w, "index.html", map[string]any{
		"WHOISRecord": whoisRecord,
	})
}
