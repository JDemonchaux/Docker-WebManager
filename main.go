package main

import (
	"log"
	"net/http"
)

const url  = "http://10.254.253.252:2375/"
//const url  = "http://192.168.0.254:2375/"
const baseUrl  = "http://127.0.0.1:1234/"

func main() {
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
	log.Println("Serving at localhost:1234...")
	log.Fatal(http.ListenAndServe(":1234", nil))
}
