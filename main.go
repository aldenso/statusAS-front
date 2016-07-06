/**
* @Author: Aldo Sotolongo
* @Date:   2016-07-03T21:02:19-04:30
* @Email:  aldenso@gmail.com
* @Last modified by:   Aldo Sotolongo
* @Last modified time: 2016-07-06T15:44:00-04:30
 */
package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	templates       *template.Template
	tomlfile        = "config.toml"
	createtemplate  bool
	apiservername   string
	apiserverport   int
	apitls          bool
	frontservername string
	frontserverport int
)

func init() {
	pageFile := "templates/default/app.html"
	templates = template.New("app.html")
	templates.Delims("{{{", "}}}")
	templates.ParseFiles(pageFile)
	flag.BoolVar(&createtemplate, "createtemplate", false, "Create an example config.toml file")
}

func readTomlFile(tomlfile string) (*Tomlconfig, error) {
	var config *Tomlconfig
	if _, err := toml.DecodeFile(tomlfile, &config); err != nil {
		return nil, err
	}
	return config, nil
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
	flag.Parse()
	if createtemplate {
		CreateTemplate()
	}
	config, err := readTomlFile(tomlfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	frontservername = config.FrontServer.FrontServerName
	frontserverport = config.FrontServer.FrontServerPort
	apiservername = config.FrontServer.APIServerName
	apiserverport = config.FrontServer.APIServerPort
	apitls = config.FrontServer.APItls
	CreateAppjs()

	r := mux.NewRouter()
	//serve static files
	r.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", http.FileServer(http.Dir("./resources/"))))
	r.HandleFunc("/dashboard", Dashboard).Methods("GET")
	loggedrouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(frontservername+":"+strconv.Itoa(frontserverport), handlers.CompressHandler(loggedrouter))
}
