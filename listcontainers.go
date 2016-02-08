package main

import (
	"io"
	"encoding/json"
)

type ListContainers []struct {
	ID string `json:"Id"`
	Names []string `json:"Names"`
	Image string `json:"Image"`
	ImageID string `json:"ImageID"`
	Command string `json:"Command"`
	Created int `json:"Created"`
	Ports []struct {
		IP string `json:"IP"`
		PrivatePort int `json:"PrivatePort"`
		PublicPort int `json:"PublicPort"`
		Type string `json:"Type"`
	} `json:"Ports"`
	Labels struct {
	   } `json:"Labels"`
	Status string `json:"Status"`
	HostConfig struct {
		   NetworkMode string `json:"NetworkMode"`
	   } `json:"HostConfig"`
	NetworkSettings struct {
		   Networks struct {
				    Bridge struct {
						   IPAMConfig interface{} `json:"IPAMConfig"`
						   Links interface{} `json:"Links"`
						   Aliases interface{} `json:"Aliases"`
						   NetworkID string `json:"NetworkID"`
						   EndpointID string `json:"EndpointID"`
						   Gateway string `json:"Gateway"`
						   IPAddress string `json:"IPAddress"`
						   IPPrefixLen int `json:"IPPrefixLen"`
						   IPv6Gateway string `json:"IPv6Gateway"`
						   GlobalIPv6Address string `json:"GlobalIPv6Address"`
						   GlobalIPv6PrefixLen int `json:"GlobalIPv6PrefixLen"`
						   MacAddress string `json:"MacAddress"`
					   } `json:"bridge"`
			    } `json:"Networks"`
	   } `json:"NetworkSettings"`
}


func (x *ListContainers)Decode(r io.Reader) (err error) {
	err = json.NewDecoder(r).Decode(x)
	return
}
