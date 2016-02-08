package main

import (
	"io"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"bytes"
)

type ListImages []struct {
	ID string `json:"Id"`
	ParentID string `json:"ParentId"`
	RepoTags []string `json:"RepoTags"`
	RepoDigests interface{} `json:"RepoDigests"`
	Created int `json:"Created"`
	Size int `json:"Size"`
	VirtualSize int `json:"VirtualSize"`
	Labels interface{} `json:"Labels"`
}

func (x *ListImages)Decode(r io.Reader) (err error) {
	err = json.NewDecoder(r).Decode(x)
	return
}

func (x *ListImages)Get() (err error){
	resp, err := http.Get(url + "images/json")
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
