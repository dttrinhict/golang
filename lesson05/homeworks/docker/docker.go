package main

import (
	"fmt"
	"github.com/dttrinhict/golang/lesson05/homeworks/docker/pattern"
	"log"
)
func main()  {
	//docCli01 := implement.NewDocker()
	//docCli02 := implement.NewDocker()
	//docCli03 := implement.NewDocker()
	//docCli04 := implement.NewDocker()
	//
	//fmt.Printf("Docker Client 01: %p,%v\n", &docCli01, docCli01)
	//fmt.Printf("Docker Client 02: %p,%v\n", &docCli02, docCli02)
	//fmt.Printf("Docker Client 03: %p,%v\n", &docCli03, docCli03)
	//fmt.Printf("Docker Client 04: %p,%v\n", &docCli04, docCli04)
	//
	//for i:=0; i <=10; i++ {
	//	fmt.Printf("Docker Client %d: %p\n",i, implement.NewDocker())
	//}

	doc :=  pattern.NewDocker()
	//fmt.Printf("Docker Client 01: %p,%v\n", &doc, doc)
	//fmt.Printf("Docker Client 02: %p,%v\n", &doc, doc)
	//err := doc.Close()
	//if err != nil {
	//	log.Printf("Create Docker error: %v", err)
	//}
	//
	//fmt.Printf("Docker Client 00: %p,%v\n", &doc, doc)

	//containers, err := doc.GetAllContainers()
	//
	//if err != nil {
	//	log.Printf("Container errorr: %v", err)
	//}
	//fmt.Printf("Docker Client 03: %p,%v\n", &doc, doc)
	//fmt.Printf("Docker Client 04: %p,%v\n", &doc, doc)
	//fmt.Println(containers)
	fmt.Println("\n\nStart Docker by ID\n")
	//ca2eed77fabbacb546d38528458891a791dff6bbb0bc8db1f1329cadde120f59
	startIDerr := doc.StartContainerByID("3611735d5512")
	if startIDerr != nil {
		log.Printf("Container errorr: %v", startIDerr)
	}

	fmt.Println("\nGet Container By Container Status\n")
	//status=(created|restarting|running|removing|paused|exited|dead)
	cons, _ := doc.GetContainersByStatus("running")
	fmt.Println(cons)

	//exciting_northcutt
	fmt.Println("\nGet container by name\n")
	cons0, err0 := doc.GetRunningContainerByName("exciting_northcutt")
	if err0 != nil {
		log.Printf("Container errorr: %v", err0)
	}
	fmt.Println(cons0)

	fmt.Println("\n\nStop Docker By Name\n")
	//ca2eed77fabbacb546d38528458891a791dff6bbb0bc8db1f1329cadde120f59
	stopNameErr := doc.StopContainerByName("exciting_northcutt")
	if stopNameErr != nil {
		log.Printf("Container errorr: %v", stopNameErr)
	}

	fmt.Println("\n\nGet container by ID\n")
	//ca2eed77fabbacb546d38528458891a791dff6bbb0bc8db1f1329cadde120f59
	cons, err1 := doc.GetRunningContainerById("3611735d5512")
	if err1 != nil {
		log.Printf("Container errorr: %v", err1)
	}
	fmt.Println(cons)
}
