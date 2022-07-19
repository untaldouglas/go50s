package main

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	//"go.uber.org/automaxprocs/maxprocs"
)

// "github.com/ardanlabs/conf"
// "github.com/dimfeld/httptreemux/v5"

var build = "develop"

func main() {

	// Set the correct number of threads for the service
	// based on what is available either by the machine or quotas assigned
	// No funcionó, pendiente de averiguar porque.
	// if _, err := maxprocs.Set(); err != nil {
	// 	fmt.Println("maxprocs: %w", err)
	// 	os.Exit(1)
	// }

	g := runtime.GOMAXPROCS(0)
	log.Printf("Starting services build[%s] CPU[%d]", build, g)
	defer log.Println("Service ended")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	log.Println("Stopping service")
}
