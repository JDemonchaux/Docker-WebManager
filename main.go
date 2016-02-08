package main

import (
	"log"
	"net/http"
)

const url  = "http://192.168.0.17:2375/"

func main() {
	DATA := new(ListContainers)
	DATA.Get()

	for _, d := range *DATA {
		log.Println(d.Image)
		in := new(InspectContainer)
		log.Println(in.Args)
		for _,i := range in.Args{
			log.Println(i)
		}
	}
	go server()
	for{

	}
}


func server(){
	http.Handle("/", http.FileServer(http.Dir("./appweb")))
	log.Println("Serving at localhost:1234...")
	log.Fatal(http.ListenAndServe(":1234", nil))
}
