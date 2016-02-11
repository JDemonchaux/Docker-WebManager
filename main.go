package main

import (
	"log"
	"net/http"
	"os"
)

//var url  = "http://10.254.253.252:2375/"
var apiUrl  = "http://192.168.0.254:2375/"

var sock = "/var/run/docker.sock"
var host,_ = os.Hostname()
var baseUrl  = "https://" + host + "/"


func main() {
	if _, err := os.Stat(sock); err != nil {
		log.Println(err)
	}else {
		go unixSock(sock)
		apiUrl  = "http://127.0.0.1:1234/"
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
	log.Println("Serving at " + host + ":443...")
	log.Fatal(http.ListenAndServeTLS(":443", "certificate/server.cert", "certificate/server.key", nil))

}
