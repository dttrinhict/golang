package main

import (
	"bufio"
	"fmt"
	"github.com/dttrinhict/golang/lesson03/homeworks/docker"
	"log"
	"os"
	"strings"
)

func main(){
	d := docker.InitDocker()
	reader := bufio.NewReader(os.Stdin)
	k := KeyInput(reader)
	f := callFunc(k, 0)
	for f != 0 {
		switch f {
		case 1:
			containersRunning, _ := d.DockerContainerListRunning()
			d.DockerPrint(containersRunning)
			k = KeyInput(reader)
			f = callFunc(k, f)
		case 2:
			containersExited, _ := d.DockerContainerListExited()
			d.DockerPrint(containersExited)
			k = KeyInput(reader)
			f = callFunc(k, f)
		case 3:
			containers, _ := d.DockerContainerListAll()
			d.DockerPrint(containers)
			k = KeyInput(reader)
			f = callFunc(k, f)
		case 4:
			containerIdOrName := strings.Trim(strings.Split(k," ")[1],"\n")
			containers, err := d.DockerGetContainerByName(containerIdOrName)
			if err != nil || containers == nil {
				if err := d.DockerContainerStart(containerIdOrName); err != nil {
					log.Printf("Error: %v", err)
				}else{
					log.Printf("The container has ID %v is started", containerIdOrName)
				}
			}else{
				for _, c := range containers {
					if err := d.DockerContainerStart(c.ID); err != nil {
						log.Printf("Error: %v", err)
					}else{
						log.Printf("The container has name %v is started", containerIdOrName)
					}
				}
			}
			k = KeyInput(reader)
			f = callFunc(k, f)
		case 5:
			containerIdOrName := strings.Trim(strings.Split(k," ")[1],"\n")
			containers, err := d.DockerGetContainerByName(containerIdOrName)
			if err != nil || containers == nil {
				if err := d.DockerContainerStop(containerIdOrName); err != nil {
					log.Printf("Error: %v", err)
				}else{
					log.Printf("The container has ID %v is stopped", containerIdOrName)
				}
			}else{
				for _, c := range containers {
					if err := d.DockerContainerStop(c.ID); err != nil {
						log.Printf("Error: %v", err)
					}else{
						log.Printf("The container has name %v is stoped", containerIdOrName)
					}
				}
			}
			k = KeyInput(reader)
			f = callFunc(k, f)
		default:
			os.Exit(0)
		}
	}
}


func KeyInput(reader *bufio.Reader) (key string) {
	fmt.Printf("Enter (n) for lisst (stop container ID/Name) for stop (start container ID/Name) for start container (another) for quit: ")
	k, err := reader.ReadString('\n')
	key = strings.Trim(k, "\n")
	if err != nil {
		fmt.Printf(err.Error())
	}
	return key
}

func callFunc(k string, result int) int {
	if strings.Compare(k, "n") == 0 {
		result++
		if result > 3 {
			result = 1
		}
	}else if strings.HasPrefix(k, "start")  {
		result = 4
	} else if strings.HasPrefix(k, "stop") {
		result = 5
	}else{
		result = 0
	}
	return result
}