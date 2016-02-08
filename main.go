package main

import (
	"log"
	"net/http"
)

const url  = "http://192.168.0.254:2375/"

func main() {
	server()
}


func server(){
	http.Handle("/app/", http.FileServer(http.Dir("./appweb")))
	//http.HandleFunc("/containers",containers)
	http.HandleFunc("/",containers)
	log.Println("Serving at localhost:1234...")
	log.Fatal(http.ListenAndServe(":1234", nil))
}
