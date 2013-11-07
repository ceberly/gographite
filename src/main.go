package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

import (
	"graphite"
	"helper"
)

func main() {

	var (
		listen  = flag.Int("l", 8080, "port to listen on.")
		net     = flag.String("n", "tcp", "network type to connect to carbon-cache server process: 'tcp', 'udp', etc.")
		addr    = flag.String("a", "127.0.0.1", "address of carbon-cache server process.")
		port    = flag.Int("p", 2003, "port of carbon-cache server process.")
		verbose = flag.Bool("v", false, "be verbose")
	)

	switch *net {
	case "tcp", "udp":
	default:
		log.Fatal("Please supply a valid network protocol (tcp or udp)")
	}

	flag.Parse()

	graphite, err := graphite.NewWithConnection(*net, fmt.Sprintf("%s:%d", *addr, *port))
	if err != nil {
		log.Fatal(err)
	}

	graphite.Verbose = *verbose

	// This should be handled by your webserver, but is super annoying when testing.
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", 404)
	})

	// host:8080/bucket/<optional subbuckets>/unix_timestamp/value
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		key, time, value, err := helper.ParseUrl(r.URL)
		if err != nil {
			log.Printf("%v  (%s)", err, r.URL.Path)
			http.Error(w, "Bad Request", 400)
			return
		}

		graphite.Send(key, time, value)
	})

	if *verbose {
		log.Printf("Starting server process on port %d", *listen)
		log.Printf("Connecting to carbon-cache process at (%s) %s:%d", *net, *addr, *port)
	}

	err = http.ListenAndServe(fmt.Sprintf(":%d", *listen), nil)
	if err != nil {
		log.Fatal(err)
	}
}
