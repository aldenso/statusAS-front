/**
* @Author: Aldo Sotolongo
* @Date:   2016-07-03T21:02:19-04:30
* @Email:  aldenso@gmail.com
* @Last modified by:   Aldo Sotolongo
* @Last modified time: 2016-07-03T21:23:59-04:30
 */
package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type page struct {
	Title string
}

var templates *template.Template

func init() {
	pageFile := "templates/default/app.html"
	templates = template.New("app.html")
	templates.Delims("{{{", "}}}")
	templates.ParseFiles(pageFile)
}

// Dashboard get services status
func Dashboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	templates.ExecuteTemplate(w, "app.html", page{Title: "StatusAS Dashboard"})
}

func main() {
	r := mux.NewRouter()
	//serve static files
	r.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", http.FileServer(http.Dir("./resources/"))))
	r.HandleFunc("/dashboard", Dashboard).Methods("GET")
	loggedrouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe("0.0.0.0:9000", loggedrouter)
}
