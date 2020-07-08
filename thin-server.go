package main

import (
	"flag"
	"log"
	"net/http"
	"runtime"
	"strconv"
)

func init() {
	// use all CPU Cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// set command-line flags and their default values
	port := flag.Int("port", 8080, "System port for running the server")
	folder := flag.String("folder", ".", "Folder to serve on the server")
	flag.Parse()

	// set the folder to serve
	fileServer := http.FileServer(http.Dir(*folder))
	http.Handle("/", fileServer)

	// start the server
	log.Println("Starting server on port", *port, "...")
	serverPort := ":" + strconv.Itoa(*port)
	log.Fatal(http.ListenAndServe(serverPort, nil))
}
