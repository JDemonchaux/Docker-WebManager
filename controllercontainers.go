package main

import (
	"text/template"
	"net/http"
	"log"
)

func containers(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	DATA := new(ListContainers)
	DATA.Get()

	tmpl, err := template.ParseFiles("appWeb/index.html")
	if err != nil {
		log.Println(err)
	}

	tmpl.ExecuteTemplate(w, "content", DATA)

	req.Body.Close()
}