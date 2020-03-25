/*
*  FILE          : main.go
*  PROJECT       : Assignment #3
*  PROGRAMMER    : Gabriel Gurgel
*  FIRST VERSION : 2020-03-24
*  DESCRIPTION   : Logoher is a GO Microservice that serves to log messages with different
*				   levels in a server. It is also gives the ability to return those logs
 */

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

// Configuration for the connection
type Config struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}

func main() {
	// Defaults if the file doesnt exists
	IP := ""
	port := "8080"

	// Read config
	configFile, err := ioutil.ReadFile("config.json")
	if err == nil {
		currentConfig := Config{}
		json.Unmarshal([]byte(configFile), &currentConfig)
		IP = currentConfig.IP
		port = currentConfig.Port
	} else {
		log.Fatal(err)
	}

	// Set up service
	svc := logopher{}

	// Make endpoints
	logInfoHandler := httptransport.NewServer(
		makeLogInfoEndpoint(svc),
		decodeLogMessageRequest,
		encodeResponse,
	)

	logfileRequestHandler := httptransport.NewServer(
		makeLogfileRequestEndpoint(svc),
		decodeLogfileRequest,
		encodeResponse,
	)

	generateIDRequestHandler := httptransport.NewServer(
		makeGenerateIDEndpoint(svc),
		decodeIDGenerationRequest,
		encodeResponse,
	)

	// Setup routers and listen
	serverAddress := IP + ":" + port
	http.Handle("/log", logInfoHandler)
	http.Handle("/file", logfileRequestHandler)
	http.Handle("/id", generateIDRequestHandler)
	log.Fatal(http.ListenAndServe(serverAddress, nil))
}
