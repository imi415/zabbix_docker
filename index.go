package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/fsouza/go-dockerclient"
)

type discoveryStruct struct {
	Data []map[string]string `json:"data"`
}

func main() {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}
	switch os.Args[1] {
	case "Info":
		info, _ := client.Info()
		switch os.Args[2] {
		case "ID":
			fmt.Println(info.ID)
		case "Images":
			fmt.Println(info.Images)
		case "Containers":
			fmt.Println(info.Containers)
		case "ContainersPaused":
			fmt.Println(info.ContainersPaused)
		case "ContainersRunning":
			fmt.Println(info.ContainersRunning)
		case "ContainersStopped":
			fmt.Println(info.ContainersStopped)
		case "ServerVersion":
			fmt.Println(info.ServerVersion)
		default:
			fmt.Println("Parameter error")
		}
	case "Container":
		msg := readCache(os.Args[2])
		if time.Now().After(msg.Read.Add(30 * time.Second)) {
			statsChannel := make(chan *docker.Stats)
			statsOption := docker.StatsOptions{ID: os.Args[2], Stats: statsChannel}
			go client.Stats(statsOption)
			msg = <-statsChannel
			writeCache(os.Args[2], msg)
		}
		switch os.Args[3] {
		case "CPU":
			usage := msg.CPUStats.CPUUsage.TotalUsage - msg.PreCPUStats.CPUUsage.TotalUsage
			sys := msg.CPUStats.SystemCPUUsage - msg.PreCPUStats.SystemCPUUsage
			fmt.Printf("%.2f\n", float32(usage/sys*100))
		case "network":
			if os.Args[4] == "rx_bytes" {
				fmt.Println(msg.Networks["eth0"].RxBytes)
			} else if os.Args[4] == "tx_bytes" {
				fmt.Println(msg.Networks["eth0"].TxBytes)
			}
		case "memory_stats":
			fmt.Println(msg.MemoryStats.Usage)
		case "pids_stats":
			fmt.Println(msg.PidsStats.Current)
		}
	case "Discovery":
		ctnList, _ := client.ListContainers(docker.ListContainersOptions{All: false})
		dscSt := new(discoveryStruct)
		for _, ctn := range ctnList {
			m := make(map[string]string)
			m["{#CONTAINER_ID}"] = ctn.ID
			m["{#CONTAINER_IMAGE}"] = ctn.Image
			m["{CONTAINER_NAME}"] = ctn.Names[0]
			m["{#CONTAINER_IP}"] = ctn.Networks.Networks["eth0"].IPAddress
			m["{#CONTAINER_STATE}"] = ctn.State
			dscSt.Data = append(dscSt.Data, m)
		}
		jsonBytes, _ := json.Marshal(dscSt)
		fmt.Println(string(jsonBytes))
	default:
		fmt.Println("Parameter error")
	}
}

func writeCache(filename string, stats *docker.Stats) bool {
	data, _ := json.Marshal(stats)
	err := ioutil.WriteFile(os.Getenv("HOME")+"/"+filename, data, 0644)
	if err != nil {
		os.Create(os.Getenv("HOME") + "/" + filename)
	}
	ioutil.WriteFile(filename, data, 0644)
	return true
}

func readCache(filename string) *docker.Stats {
	stats := docker.Stats{Read: time.Now()}
	data, err := ioutil.ReadFile(os.Getenv("HOME") + "/" + filename)
	if err != nil {
		stats.Read = time.Unix(0, 0)
		return &stats
	}
	json.Unmarshal(data, &stats)
	return &stats
}
