package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/fsouza/go-dockerclient"
)

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
		case "ContainerPaused":
			fmt.Println(info.ContainersPaused)
		case "ContainerRunning":
			fmt.Println(info.ContainersRunning)
		case "ContainerStopped":
			fmt.Println(info.ContainersRunning)
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
	default:
		fmt.Println("Parameter error")
	}
}

func writeCache(filename string, stats *docker.Stats) bool {
	data, _ := json.Marshal(stats)
	err := ioutil.WriteFile(os.Getenv("HOME")+"/"+filename, data, 0644)
	if err != nil {
		os.Create(filename)
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
