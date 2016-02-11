package main

import (
	"io"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"bytes"
)

type ToptContainer struct {
	Processes *[][]string `json:"Processes"`
	Titles []string `json:"Titles"`
}


func (x *ToptContainer)Decode(r io.Reader) (err error) {
	err = json.NewDecoder(r).Decode(x)
	return
}

func (x *ToptContainer)Get(ID string) (err error){
	resp, err := http.Get(apiUrl + "containers/"+ID+"/top?ps_args=aux")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	x.Decode(bytes.NewReader(body))

	resp.Body.Close()
	return
}