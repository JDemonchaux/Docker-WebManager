package main

import (
	"time"
	"io"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"bytes"
)

type InspectImages struct {
	ID string `json:"Id"`
	RepoTags []string `json:"RepoTags"`
	RepoDigests []interface{} `json:"RepoDigests"`
	Parent string `json:"Parent"`
	Comment string `json:"Comment"`
	Created time.Time `json:"Created"`
	Container string `json:"Container"`
	ContainerConfig struct {
		   Hostname string `json:"Hostname"`
		   Domainname string `json:"Domainname"`
		   User string `json:"User"`
		   AttachStdin bool `json:"AttachStdin"`
		   AttachStdout bool `json:"AttachStdout"`
		   AttachStderr bool `json:"AttachStderr"`
		   Tty bool `json:"Tty"`
		   OpenStdin bool `json:"OpenStdin"`
		   StdinOnce bool `json:"StdinOnce"`
		   Env interface{} `json:"Env"`
		   Cmd []string `json:"Cmd"`
		   Image string `json:"Image"`
		   Volumes interface{} `json:"Volumes"`
		   WorkingDir string `json:"WorkingDir"`
		   Entrypoint interface{} `json:"Entrypoint"`
		   OnBuild interface{} `json:"OnBuild"`
		   Labels interface{} `json:"Labels"`
	   } `json:"ContainerConfig"`
	DockerVersion string `json:"DockerVersion"`
	Author string `json:"Author"`
	Config struct {
		   Hostname string `json:"Hostname"`
		   Domainname string `json:"Domainname"`
		   User string `json:"User"`
		   AttachStdin bool `json:"AttachStdin"`
		   AttachStdout bool `json:"AttachStdout"`
		   AttachStderr bool `json:"AttachStderr"`
		   Tty bool `json:"Tty"`
		   OpenStdin bool `json:"OpenStdin"`
		   StdinOnce bool `json:"StdinOnce"`
		   Env interface{} `json:"Env"`
		   Cmd []string `json:"Cmd"`
		   Image string `json:"Image"`
		   Volumes interface{} `json:"Volumes"`
		   WorkingDir string `json:"WorkingDir"`
		   Entrypoint interface{} `json:"Entrypoint"`
		   OnBuild interface{} `json:"OnBuild"`
		   Labels interface{} `json:"Labels"`
	   } `json:"Config"`
	Architecture string `json:"Architecture"`
	Os string `json:"Os"`
	Size int `json:"Size"`
	VirtualSize int `json:"VirtualSize"`
	GraphDriver struct {
		   Name string `json:"Name"`
		   Data interface{} `json:"Data"`
	   } `json:"GraphDriver"`
}

func (x *InspectImages)Decode(r io.Reader) (err error) {
	err = json.NewDecoder(r).Decode(x)
	return
}

func (x *InspectImages)Get(ID string) (err error){
	resp, err := http.Get(url + "images/"+ID+"/json")
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
