package main

import (
	"log"
	"net/http"
	"os"
	"./settings"
	"./authentication"
)

var config settings.SettingsType

func main() {


	err := config.ReadSettings()
	if err != nil {
		config.ApiUrl = "http://192.168.0.254:2375/"
		config.BaseUrl = "http://127.0.0.1:8080/"
		config.Sock = "/var/run/docker.sock"
		config.SavSettings()
	}

	if _, err := os.Stat(config.Sock); err != nil {
		log.Println(err)
	}else {
		config.ApiUrl  = "http://127.0.0.1:1234/"
		go unixSock(config.Sock)
	}
	server()

}


func server(){
	http.Handle("/app/", http.FileServer(http.Dir("./appweb")))
	
	//http.HandleFunc("/containers",containers)
	http.HandleFunc("/",containers)
	//test login
	login := new(authentication.AuthenticationType)
	http.HandleFunc("/login", login.Auth)

	//route
	http.HandleFunc("/containers/rename/", containersRename)
	http.HandleFunc("/containers/inspect/", containersInspect)
	http.HandleFunc("/containers/restart/", containersRestart)
	http.HandleFunc("/containers/start/", containersStart)
	http.HandleFunc("/containers/stop/", containersStop)
	http.HandleFunc("/containers/pause/", containersPause)
	http.HandleFunc("/containers/unpause/", containersUnpause)
	http.HandleFunc("/containers/delete/", containersDelete)
	log.Println("Serving at " + config.BaseUrl + "...")
	log.Fatal(http.ListenAndServe(":8080", nil))
	//log.Println("Serving at " + config.BaseUrl + ":443...")
	//log.Fatal(http.ListenAndServeTLS(":443", "certificate/server.cert", "certificate/server.key", nil))

}
