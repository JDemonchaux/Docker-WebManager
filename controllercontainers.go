package main

import (
	"text/template"
	"net/http"
	"log"
	"regexp"
	"strings"
	"bytes"
	"strconv"
)

func containersStart(w http.ResponseWriter, req *http.Request) {
	p := strings.Split(req.URL.Path, "/")

	_, err := http.Post(url + "containers/" + p[len(p) - 1] + "/start", "", nil)
	if err != nil {
		log.Println(err)
	}

	log.Println(http.StatusFound)
	http.Redirect(w, req, baseUrl, http.StatusFound)
	return
}
func containersStop(w http.ResponseWriter, req *http.Request) {
	p := strings.Split(req.URL.Path, "/")

	_, err := http.Post(url + "containers/" + p[len(p) - 1] + "/stop", "", nil)
	if err != nil {
		log.Println(err)
	}

	log.Println(http.StatusFound)
	http.Redirect(w, req, baseUrl, http.StatusFound)
	return
}
func containersPause(w http.ResponseWriter, req *http.Request) {
	p := strings.Split(req.URL.Path, "/")

	_, err := http.Post(url + "containers/" + p[len(p) - 1] + "/pause", "", nil)
	if err != nil {
		log.Println(err)
	}

	log.Println(http.StatusFound)
	http.Redirect(w, req, baseUrl, http.StatusFound)
	return
}
func containersRestart(w http.ResponseWriter, req *http.Request) {
	p := strings.Split(req.URL.Path, "/")

	_, err := http.Post(url + "containers/" + p[len(p) - 1] + "/restart", "", nil)
	if err != nil {
		log.Println(err)
	}

	log.Println(http.StatusFound)
	http.Redirect(w, req, baseUrl, http.StatusFound)
	return
}
func containersUnpause(w http.ResponseWriter, req *http.Request) {
	p := strings.Split(req.URL.Path, "/")

	_, err := http.Post(url + "containers/" + p[len(p) - 1] + "/unpause", "", nil)
	if err != nil {
		log.Println(err)
	}

	log.Println(http.StatusFound)
	http.Redirect(w, req, baseUrl, http.StatusFound)
	return
}

func containersDelete(w http.ResponseWriter, req *http.Request) {
	p := strings.Split(req.URL.Path, "/")

	client := &http.Client{}
	reqDelete, err := http.NewRequest(
		"DELETE",
		url + "containers/" + p[len(p) - 1] + "?v=1",
		bytes.NewBuffer([]byte("[]")))
	if err != nil {
		log.Println(err)
	}
	client.Do(reqDelete)

	log.Println(http.StatusFound)
	http.Redirect(w, req, baseUrl, http.StatusFound)
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

	tmpl, err := template.ParseFiles("appWeb/header.html", "appWeb/index.html", "appWeb/footer.html")

	if err != nil {
		log.Println(err)
	}

	tmpl.ExecuteTemplate(w, "header", nil)
	tmpl.ExecuteTemplate(w, "index", data)
	tmpl.ExecuteTemplate(w, "footer", nil)
	req.Body.Close()
}

func containersInspect(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	log.Println("containersInspect")
	p := strings.Split(req.URL.Path, "/")
	id := p[len(p) - 1]
	log.Println("index")

	ic := new(InspectContainer)

	ic.Get(id)
	ii := new(InspectImages)
	ii.Get(ic.Image)
	tc := new(ToptContainer)
	tc.Get(ic.ID)
	//for _, t := range *tc.Processes {
	//	for i, u := range t {
	//		log.Println(tc.Titles[i] + ": " + u)
	//	}
	//}
	contJson := new(ListContainers)
	contJson.GetByID(id)

	type ContainerInfo struct {
		ID         string
		Name       string
		RepoTags   []string
		Args       []string
		Path       string
		Hostname   string
		Domainname string
		Env        []string
		ShmSize    int
	}
	type ExposedPorts struct {
		Private string
		Public  string
	}
	type NetworkInfo struct {
		IPAddress    string
		Gateway      string
		MacAddress   string
		ExposedPorts []ExposedPorts
	}
	type DATA struct {
		ContainerInfo
		NetworkInfo
	}

	contInfo := new(ContainerInfo)
	contInfo.ID = ic.ID
	contInfo.Domainname = ic.Config.Domainname
	contInfo.Name = ic.Name
	contInfo.RepoTags = ii.RepoTags
	contInfo.Args = ic.Args
	contInfo.Path = ic.Path
	contInfo.Hostname = ic.Config.Hostname
	contInfo.Env = ic.Config.Env
	contInfo.ShmSize = ic.HostConfig.ShmSize

	netInfo := new(NetworkInfo)
	netInfo.IPAddress = ic.NetworkSettings.IPAddress
	netInfo.Gateway = ic.NetworkSettings.Gateway
	netInfo.MacAddress = ic.NetworkSettings.MacAddress
	log.Println(contJson)
	for _, c := range *contJson{
		for _,port := range c.Ports {
			netInfo.ExposedPorts = append(netInfo.ExposedPorts,
				ExposedPorts{strconv.FormatInt(int64(port.PrivatePort),10),
					strconv.FormatInt(int64(port.PublicPort),10)+"/"+port.Type})
		}
	}
	log.Println(netInfo.ExposedPorts)
	data := new(DATA)
	data.ContainerInfo = *contInfo
	data.NetworkInfo = *netInfo

	tmpl, err := template.ParseFiles("appWeb/header.html", "appWeb/container-detail.html", "appWeb/footer.html")

	if err != nil {
		log.Println(err)
	}

	tmpl.ExecuteTemplate(w, "header", nil)
	tmpl.ExecuteTemplate(w, "index", data)
	tmpl.ExecuteTemplate(w, "footer", nil)
	req.Body.Close()
}