package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	flagEndpoint string
	flagPort     int
)

func main() {
	logger := log.New(os.Stdout, "health-check ", log.Flags())

	// allow flags to configure the container
	flag.StringVar(&flagEndpoint, "endpoint", "/", "The endpoint to listen for dummy healthchecks")
	flag.IntVar(&flagPort, "port", 8080, "The port the dummy healthcheck should respond on")
	flag.Parse()

	// check if the environment variable is defined for each and override
	endpoint, v := os.LookupEnv("APP_ENDPOINT")
	if v {
		flagEndpoint = endpoint
	}

	port, v := os.LookupEnv("APP_PORT")
	if v {
		// convert to an int, if it errors leave as default
		if p, err := strconv.Atoi(port); err == nil {
			flagPort = p
		}
	}

	logger.Println("listening on port", flagPort)
	logger.Println("accepting traffic on", flagEndpoint)

	// handle the healthcheck endpoint
	http.HandleFunc(flagEndpoint, func(w http.ResponseWriter, r *http.Request) {
		logger.Println("Host:", r.Host, "Method:", r.Method, "URI:", r.RequestURI, "Headers:", r.Header)
		fmt.Fprintf(w, "Healthy")
	})

	// listen to port
	http.ListenAndServe(fmt.Sprintf(":%d", flagPort), nil)
}
