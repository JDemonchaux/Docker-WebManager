package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"bytes"
)

func main() {
	resp, err := http.Get("http://10.0.50.96:2375/containers/json?all=1")
	if err != nil {
		log.Panicln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	DATA := new(ContainersJsonArr)

	DATA.Decode(bytes.NewReader(body))
	for _, d := range DATA {
		log.Println(d.ID)
	}
	resp.Body.Close()

}

