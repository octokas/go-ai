package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/octokas/go-ai/pkg/router"
)

func HelloDutonian() string {
	return "Hello, Dutonian! 👋🏼"
}

func main() {

	// Define flags for different servers
	apiV1 := flag.Bool("apiv1", false, "Start only APIv1 server")
	apiV2 := flag.Bool("apiv2", false, "Start only APIv2 server")
	home := flag.Bool("home", false, "Start only home server")
	chat := flag.Bool("chat", false, "Start only chat server")
	flag.Parse()

	// Channel to collect errors from all components
	errorChan := make(chan error, 10)
	//done := make(chan bool)

	// Error monitoring goroutine
	go func() {
		for err := range errorChan {
			log.Printf("ERROR: %+v\n", err)
		}
	}()

	fmt.Println(HelloDutonian())

	// Setup routes based on flags
	if !*apiV2 && !*home && !*apiV1 && !*chat {
		log.Println("Starting all servers...")
		go router.HomeServer() // 8080
		go router.V1Server()   // 2020
		go router.V2Server()   // 3030
		go router.ChatServer() // 4040
	} else if *apiV2 {
		go router.V2Server()
	} else if *home {
		go router.HomeServer()
	} else if *apiV1 {
		go router.V1Server()
	} else if *chat {
		go router.ChatServer()
	} else {
		log.Println("No server type specified. Use -all flag to start all servers")
		return
	}

	// Block forever to keep the program running
	select {}

}
