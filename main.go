package main

import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	svc := logopher{}

	logInfoHandler := httptransport.NewServer(
		makeLogInfoEndpoint(svc),
		decodeLogMessageRequest,
		encodeResponse,
	)

	http.Handle("/logInfo", logInfoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
