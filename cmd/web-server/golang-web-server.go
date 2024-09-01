package main

import (
	"flag"
	"log"
	"web_server/configs"
	"web_server/pkg/mhttp"
)

var useHandlerFunctions bool

func main() {
	flag.BoolVar(&useHandlerFunctions, "use-handler-functions", true, "Runs the web server with handler functions, use false to run it in fileserver mode.")
	flag.Parse()

	server := mhttp.NewServer(configs.SERVER_STATIC_DIRECTORY, configs.SERVER_URL, configs.SERVER_PORT)
	var err error

	if useHandlerFunctions {
		err = server.InitializeHandlerFunctions()
	} else {
		err = server.InitializeFileServer()
	}

	if err != nil {
		log.Fatal(err)
	}
	server.ListenAndServe()
}
