package main

import (
	"fmt"
	"github.com/dttrinhict/golang/lesson03/homeworks/calendar"
	"github.com/dttrinhict/golang/lesson03/homeworks/docker"
)

func main()  {
	_ = calendar.Calendar()
	//calendar.PrintCalendar(cal)

	d := docker.InitDocker()
	if containers, err := d.DockerContainerListAll(); err != nil {
		panic(err)
	}else{
		for _, c := range containers {
			fmt.Printf("%6s %30s  %d %10s %30s\n",c.ID, c.Names[0], c.Created, c.State, c.Status)
			//fmt.Printf("%s %s %d %s %s\n",c.ID, c.Names[0], c.Created, c.State, c.Status)
		}
	}

	fmt.Println()
	containersRunning, _ := d.DockerContainerListRunning()
	for _, c := range containersRunning {
		fmt.Printf("%6s %30s  %d %10s %30s\n",c.ID, c.Names[0], c.Created, c.State, c.Status)
		//fmt.Printf("%s %s %d %s %s\n",c.ID, c.Names[0], c.Created, c.State, c.Status)
	}

	fmt.Println()
	containersExited, _ := d.DockerContainerListExited()
	for _, c := range containersExited {
		fmt.Printf("%6s %30s  %d %10s %30s\n",c.ID, c.Names[0], c.Created, c.State, c.Status)
		//fmt.Printf("%s %s %d %s %s\n",c.ID, c.Names[0], c.Created, c.State, c.Status)
	}
}
