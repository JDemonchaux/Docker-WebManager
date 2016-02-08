package main

import (
	"text/template"
	"net/http"
	"log"
)

func containers(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	log.Println("index")
	DATA := new(ListContainers)
	DATA.Get()

	tmpl, err := template.ParseFiles("appWeb/index.html")
	if err != nil {
		log.Println(err)
	}

	tmpl.ExecuteTemplate(w, "index", DATA)

	req.Body.Close()
}