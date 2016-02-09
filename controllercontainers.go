package main

import (
	"text/template"
	"net/http"
	"log"
	"regexp"
	"strings"
)

func containersStart(w http.ResponseWriter, req *http.Request) {
	p := strings.Split(req.URL.Path, "/")

	_, err := http.Post(url + "containers/" + p[len(p) - 1] + "/start", "", nil)
	if err != nil {
		log.Println(err)
	}

	log.Println(http.StatusFound)
	http.Redirect(w,req,baseUrl,http.StatusFound)
	return
}
func containersStop(w http.ResponseWriter, req *http.Request) {
	p := strings.Split(req.URL.Path, "/")

	_, err := http.Post(url + "containers/" + p[len(p) - 1] + "/stop", "", nil)
	if err != nil {
		log.Println(err)
	}

	log.Println(http.StatusFound)
	http.Redirect(w,req,baseUrl,http.StatusFound)
	return
}
func containersPause(w http.ResponseWriter, req *http.Request) {
	p := strings.Split(req.URL.Path, "/")

	_, err := http.Post(url + "containers/" + p[len(p) - 1] + "/pause", "", nil)
	if err != nil {
		log.Println(err)
	}

	log.Println(http.StatusFound)
	http.Redirect(w,req,baseUrl,http.StatusFound)
	return
}
func containersUnpause(w http.ResponseWriter, req *http.Request) {
	p := strings.Split(req.URL.Path, "/")

	_, err := http.Post(url + "containers/" + p[len(p) - 1] + "/unpause", "", nil)
	if err != nil {
		log.Println(err)
	}

	log.Println(http.StatusFound)
	http.Redirect(w,req,baseUrl,http.StatusFound)
	return
}

func containers(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	log.Println("index")
	lc := new(ListContainers)
	lc.Get()

	type DATA []struct {
		ID         string
		Status     string
		StatusView string
		Names      []string
		Image      string
	}
	data := make(DATA, len(*lc))

	for i, d := range *lc {
		data[i].ID = d.ID
		data[i].Status = d.Status
		data[i].Names = d.Names
		data[i].Image = d.Image
		if b, _ := regexp.MatchString("Up", d.Status); b != false {
			data[i].StatusView = "container-running"
		}
		if b, _ := regexp.MatchString("Paused", d.Status); b != false {
			data[i].StatusView = "container-paused"
		}
		if b, _ := regexp.MatchString("Exited", d.Status); b != false {
			data[i].StatusView = "container-stopped"
		}
	}

	tmpl, err := template.ParseFiles("appWeb/header.html","appWeb/index.html","appWeb/footer.html")

	if err != nil {
		log.Println(err)
	}

	tmpl.ExecuteTemplate(w, "header", nil)
	tmpl.ExecuteTemplate(w, "index", data)
	tmpl.ExecuteTemplate(w, "footer", nil)
	req.Body.Close()
}

func containersInspect (w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("appWeb/header.html","appWeb/container-detail.html","appWeb/footer.html")

	if err != nil {
		log.Println(err)
	}

	tmpl.ExecuteTemplate(w, "header", nil)
	tmpl.ExecuteTemplate(w, "index", nil)
	tmpl.ExecuteTemplate(w, "footer", nil)
	req.Body.Close()
}