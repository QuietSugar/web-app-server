package main

import (
	"flag"
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"log"
	"net"
	"net/http"
	"path/filepath"
	"time"
)

var (
	port = flag.Int("p", getAnAvailablePort(), "Port of the server")
	root = flag.String("w", ".", "Root dir of the server")
)

func init() {
	flag.Parse()
}

var url string

func main() {
	a, err := filepath.Abs(*root)
	if err != nil {
		log.Fatal(err)
	}
	d := http.Dir(a)
	fs := http.FileServer(d)
	http.Handle("/", fs)
	log.Printf("Web Server on :%d...\n", *port)
	log.Printf("Web Server path :%s...\n", a)
	go func() {
		log.Println(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
	}()
	time.Sleep(time.Second)
	url = fmt.Sprintf("http://localhost:%d", *port)
	open.Run(url)
	startTray()
}

func portCheck(port int) bool {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return false
	}
	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
			fmt.Println("Failed to Close Listener :", err)
		}
	}(l)
	return true
}

func getAnAvailablePort() int {
	startPort := 8080
	endPort := 9080
	for port := startPort; port <= endPort; port++ {
		if portCheck(port) {
			return port
		}
	}
	return startPort
}
