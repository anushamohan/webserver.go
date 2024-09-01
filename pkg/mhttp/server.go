package mhttp

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

type IServer interface {
	InitializeFileServer() error
	InitializeHandlerFunctions() error
	ListenAndServe()
}

type ServerConfig struct {
	staticDir string
	url       string
	port      string
}

func NewServer(staticDir string, url string, port string) ServerConfig {
	var serverConfig = ServerConfig{
		staticDir: staticDir,
		url:       url,
		port:      port,
	}
	return serverConfig
}

func (fs ServerConfig) ListenAndServe() {
	serverUrl := fs.url + ":" + fs.port
	fmt.Println("Listening on %s...\n", serverUrl)
	log.Fatal(http.ListenAndServe(serverUrl, nil))
}

func (fs ServerConfig) InitializeFileServer() error {
	fmt.Println("Initializing handler with file server")

	if err := validateFolder(fs.staticDir); err != nil {
		return errors.New("couldn't validate static folder, error: " + err.Error())
	}

	http.Handle("/", http.FileServer(http.Dir(fs.staticDir)))
	fmt.Println("Handler initialized")
	return nil

}

func (fs ServerConfig) InitializeHandlerFunctions() error {
	staticDir = fs.staticDir
	fmt.Println("Initializing handler with functions")
	if err := validateFolder(fs.staticDir); err != nil {
		return errors.New("couldn't validate static folder, error: " + err.Error())
	}
	http.HandleFunc("/", serveFile)
	http.HandleFunc("/increment", incrementCounter)
	http.HandleFunc("/hi", hiString)
	fmt.Println("Handler functions have been initialized")
	return nil

}

func validateFolder(folderpath string) error {
	fileInfo, err := os.Stat(folderpath)
	if err != nil {
		return err
	}
	if fileInfo != nil && !fileInfo.IsDir() {
		return fmt.Errorf("%s : is not a directory", folderpath)
	}
	return nil
}
