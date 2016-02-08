package main

import (
	"text/template"
	"net/http"
	"log"
	"regexp"
)

func containers2(w http.ResponseWriter, req *http.Request) {
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

func containers(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	log.Println("index")
	lc := new(ListContainers)
	lc.Get()

	type DATA []struct {
		ID string
		Status string
		StatusView string
		Names []string
		Image string
	}
	data := make(DATA, len(*lc))

	for i,d := range *lc {
		data[i].ID = d.ID
		data[i].Status = d.Status
		data[i].Names = d.Names
		data[i].Image = d.Image
		if b,_ := regexp.MatchString("Up", d.Status); b != false {
			data[i].StatusView = "container-running"
		}
		if b,_ := regexp.MatchString("Paused", d.Status); b != false {
			data[i].StatusView = "container-paused"
		}
		if b,_ := regexp.MatchString("Exited", d.Status); b != false {
			data[i].StatusView = "container-stopped"
		}
	}
	for i := range data {
		log.Println(data[i].StatusView)
	}

	tmpl, err := template.ParseFiles("appWeb/index.html")
	if err != nil {
		log.Println(err)
	}

	tmpl.ExecuteTemplate(w, "index", data)

	req.Body.Close()
}