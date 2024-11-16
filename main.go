package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/octokas/go-ai/pkg/api"
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
	//all := flag.Bool("all", false, "Start all servers")
	//port := flag.String("port", "8080", "Server port")
	flag.Parse()

	// Channel to collect errors from all components
	errorChan := make(chan error, 10)

	// Error monitoring goroutine
	go func() {
		for err := range errorChan {
			log.Printf("ERROR: %+v\n", err) // %+v for detailed error output
		}
	}()

	// Setup all routes
	//router.SetupAll()

	fmt.Println(HelloDutonian())

	// Setup routes based on flags
	if !*apiV2 && !*home {
		log.Println("Starting all servers...")
		go router.SetupAll()

		// Start the server when using flags
		// log.Printf("Server starting on port %s", *port)
		// if err := http.ListenAndServe(fmt.Sprintf(":%s", *port), nil); err != nil {
		// 	errorChan <- err
		// }
	} else if *apiV2 {
		//log.Println("Starting API v2 server...")
		go api.RunAPIv2()
	} else if *home {
		//log.Println("Starting home server...")
		go router.HomeServer()
	} else if *apiV1 {
		//log.Println("Starting API v1 server...")
		go router.V1Server()
	} else {
		log.Println("No server type specified. Use -all flag to start all servers")
		return
	}

	// Block forever to keep the program running
	select {}

	// Start the server
	// log.Println("Server starting on :8080")
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	errorChan <- err
	// }

}
