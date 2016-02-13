package main

import (
	"log"
	"net/http"
	"os"
)

var settings SettingType


func main() {


	err := settings.ReadSettings()
	if err != nil {
		settings.ApiUrl = "http://192.168.0.254:2375/"
		settings.BaseUrl = "https://127.0.0.1/"
		settings.Sock = "/var/run/docker.sock"
		settings.SavSeettings()
	}

	if _, err := os.Stat(settings.Sock); err != nil {
		log.Println(err)
	}else {
		settings.ApiUrl  = "http://127.0.0.1:1234/"
		go unixSock(settings.Sock)
	}
	server()

}


func server(){
	http.Handle("/app/", http.FileServer(http.Dir("./appweb")))
	//http.HandleFunc("/containers",containers)
	http.HandleFunc("/",containers)

	http.HandleFunc("/containers/inspect/", containersInspect)
	http.HandleFunc("/containers/restart/", containersRestart)
	http.HandleFunc("/containers/start/", containersStart)
	http.HandleFunc("/containers/stop/", containersStop)
	http.HandleFunc("/containers/pause/", containersPause)
	http.HandleFunc("/containers/unpause/", containersUnpause)
	http.HandleFunc("/containers/delete/", containersDelete)
	log.Println("Serving at " + settings.BaseUrl + ":443...")
	log.Fatal(http.ListenAndServeTLS(":443", "certificate/server.cert", "certificate/server.key", nil))

}
