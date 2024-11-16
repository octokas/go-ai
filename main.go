package main

import (
	"fmt"
	"log"
)

func HelloDutonian() string {
	return "Hello, Dutonian! 👋🏼"
}

func main() {

	// Channel to collect errors from all components
	errorChan := make(chan error, 10)

	// Error monitoring goroutine
	go func() {
		for err := range errorChan {
			log.Printf("ERROR: %+v\n", err) // %+v for detailed error output
		}
	}()

	fmt.Println(HelloDutonian())

}
