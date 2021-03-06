package main

import (
	"time"
	"io"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"bytes"
)

type InspectContainer struct {
	ID string `json:"Id"`
	Created time.Time `json:"Created"`
	Path string `json:"Path"`
	Args []string `json:"Args"`
	State struct {
		   Status string `json:"Status"`
		   Running bool `json:"Running"`
		   Paused bool `json:"Paused"`
		   Restarting bool `json:"Restarting"`
		   OOMKilled bool `json:"OOMKilled"`
		   Dead bool `json:"Dead"`
		   Pid int `json:"Pid"`
		   ExitCode int `json:"ExitCode"`
		   Error string `json:"Error"`
		   StartedAt time.Time `json:"StartedAt"`
		   FinishedAt time.Time `json:"FinishedAt"`
	   } `json:"State"`
	Image string `json:"Image"`
	ResolvConfPath string `json:"ResolvConfPath"`
	HostnamePath string `json:"HostnamePath"`
	HostsPath string `json:"HostsPath"`
	LogPath string `json:"LogPath"`
	Name string `json:"Name"`
	RestartCount int `json:"RestartCount"`
	Driver string `json:"Driver"`
	MountLabel string `json:"MountLabel"`
	ProcessLabel string `json:"ProcessLabel"`
	AppArmorProfile string `json:"AppArmorProfile"`
	ExecIDs interface{} `json:"ExecIDs"`
	HostConfig struct {
		   Binds interface{} `json:"Binds"`
		   ContainerIDFile string `json:"ContainerIDFile"`
		   LogConfig struct {
				 Type string `json:"Type"`
				 Config struct {
				      } `json:"Config"`
			 } `json:"LogConfig"`
		   NetworkMode string `json:"NetworkMode"`
		   PortBindings struct {
				 Three306TCP []struct {
					 HostIP string `json:"HostIp"`
					 HostPort string `json:"HostPort"`
				 } `json:"3306/tcp"`
				 Eight0TCP []struct {
					 HostIP string `json:"HostIp"`
					 HostPort string `json:"HostPort"`
				 } `json:"80/tcp"`
			 } `json:"PortBindings"`
		   RestartPolicy struct {
				 Name string `json:"Name"`
				 MaximumRetryCount int `json:"MaximumRetryCount"`
			 } `json:"RestartPolicy"`
		   VolumeDriver string `json:"VolumeDriver"`
		   VolumesFrom interface{} `json:"VolumesFrom"`
		   CapAdd interface{} `json:"CapAdd"`
		   CapDrop interface{} `json:"CapDrop"`
		   DNS []interface{} `json:"Dns"`
		   DNSOptions []interface{} `json:"DnsOptions"`
		   DNSSearch []interface{} `json:"DnsSearch"`
		   ExtraHosts interface{} `json:"ExtraHosts"`
		   GroupAdd interface{} `json:"GroupAdd"`
		   IpcMode string `json:"IpcMode"`
		   Links interface{} `json:"Links"`
		   OomScoreAdj int `json:"OomScoreAdj"`
		   PidMode string `json:"PidMode"`
		   Privileged bool `json:"Privileged"`
		   PublishAllPorts bool `json:"PublishAllPorts"`
		   ReadonlyRootfs bool `json:"ReadonlyRootfs"`
		   SecurityOpt interface{} `json:"SecurityOpt"`
		   UTSMode string `json:"UTSMode"`
		   ShmSize int `json:"ShmSize"`
		   ConsoleSize []int `json:"ConsoleSize"`
		   Isolation string `json:"Isolation"`
		   CPUShares int `json:"CpuShares"`
		   CgroupParent string `json:"CgroupParent"`
		   BlkioWeight int `json:"BlkioWeight"`
		   BlkioWeightDevice interface{} `json:"BlkioWeightDevice"`
		   BlkioDeviceReadBps interface{} `json:"BlkioDeviceReadBps"`
		   BlkioDeviceWriteBps interface{} `json:"BlkioDeviceWriteBps"`
		   BlkioDeviceReadIOps interface{} `json:"BlkioDeviceReadIOps"`
		   BlkioDeviceWriteIOps interface{} `json:"BlkioDeviceWriteIOps"`
		   CPUPeriod int `json:"CpuPeriod"`
		   CPUQuota int `json:"CpuQuota"`
		   CpusetCpus string `json:"CpusetCpus"`
		   CpusetMems string `json:"CpusetMems"`
		   Devices []interface{} `json:"Devices"`
		   KernelMemory int `json:"KernelMemory"`
		   Memory int `json:"Memory"`
		   MemoryReservation int `json:"MemoryReservation"`
		   MemorySwap int `json:"MemorySwap"`
		   MemorySwappiness int `json:"MemorySwappiness"`
		   OomKillDisable bool `json:"OomKillDisable"`
		   PidsLimit int `json:"PidsLimit"`
		   Ulimits interface{} `json:"Ulimits"`
	   } `json:"HostConfig"`
	GraphDriver struct {
		   Name string `json:"Name"`
		   Data interface{} `json:"Data"`
	   } `json:"GraphDriver"`
	SizeRw int `json:"SizeRw"`
	SizeRootFs int `json:"SizeRootFs"`
	Mounts []struct {
		Name string `json:"Name"`
		Source string `json:"Source"`
		Destination string `json:"Destination"`
		Driver string `json:"Driver"`
		Mode string `json:"Mode"`
		RW bool `json:"RW"`
		Propagation string `json:"Propagation"`
	} `json:"Mounts"`
	Config struct {
		   Hostname string `json:"Hostname"`
		   Domainname string `json:"Domainname"`
		   User string `json:"User"`
		   AttachStdin bool `json:"AttachStdin"`
		   AttachStdout bool `json:"AttachStdout"`
		   AttachStderr bool `json:"AttachStderr"`
		   ExposedPorts struct {
				    Three306TCP struct {
						} `json:"3306/tcp"`
				    Eight0TCP struct {
						} `json:"80/tcp"`
			    } `json:"ExposedPorts"`
		   Tty bool `json:"Tty"`
		   OpenStdin bool `json:"OpenStdin"`
		   StdinOnce bool `json:"StdinOnce"`
		   Env []string `json:"Env"`
		   Cmd []string `json:"Cmd"`
		   Image string `json:"Image"`
		   Volumes struct {
				    EtcMysql struct {
					     } `json:"/etc/mysql"`
				    VarLibMysql struct {
					     } `json:"/var/lib/mysql"`
			    } `json:"Volumes"`
		   WorkingDir string `json:"WorkingDir"`
		   Entrypoint interface{} `json:"Entrypoint"`
		   OnBuild interface{} `json:"OnBuild"`
		   Labels struct {
			    } `json:"Labels"`
		   StopSignal string `json:"StopSignal"`
	   } `json:"Config"`
	NetworkSettings struct {
		   Bridge string `json:"Bridge"`
		   SandboxID string `json:"SandboxID"`
		   HairpinMode bool `json:"HairpinMode"`
		   LinkLocalIPv6Address string `json:"LinkLocalIPv6Address"`
		   LinkLocalIPv6PrefixLen int `json:"LinkLocalIPv6PrefixLen"`
		   Ports struct {
				  Three306TCP []struct {
					  HostIP string `json:"HostIp"`
					  HostPort string `json:"HostPort"`
				  } `json:"3306/tcp"`
				  Eight0TCP []struct {
					  HostIP string `json:"HostIp"`
					  HostPort string `json:"HostPort"`
				  } `json:"80/tcp"`
			  } `json:"Ports"`
		   SandboxKey string `json:"SandboxKey"`
		   SecondaryIPAddresses interface{} `json:"SecondaryIPAddresses"`
		   SecondaryIPv6Addresses interface{} `json:"SecondaryIPv6Addresses"`
		   EndpointID string `json:"EndpointID"`
		   Gateway string `json:"Gateway"`
		   GlobalIPv6Address string `json:"GlobalIPv6Address"`
		   GlobalIPv6PrefixLen int `json:"GlobalIPv6PrefixLen"`
		   IPAddress string `json:"IPAddress"`
		   IPPrefixLen int `json:"IPPrefixLen"`
		   IPv6Gateway string `json:"IPv6Gateway"`
		   MacAddress string `json:"MacAddress"`
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
	//custom
	RawData string
}

func (x *InspectContainer)Decode(r io.Reader) (err error) {
	err = json.NewDecoder(r).Decode(x)
	return
}

func (x *InspectContainer)Get(ID string) (err error){
	resp, err := http.Get(config.ApiUrl + "containers/"+ID+"/json?size=1")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	x.Decode(bytes.NewReader(body))
	x.RawData = string(body)

	resp.Body.Close()
	return
}