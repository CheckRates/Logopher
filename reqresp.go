/*
*  FILE          : reqresp.go
*  PROJECT       : Assignment #3
*  PROGRAMMER    : Gabriel Gurgel
*  FIRST VERSION : 2020-03-24
*  DESCRIPTION   : Defines all the responses and requests necessary for communcation
*				between client and service
 */
package main

import (
	"context"
	"encoding/json"
	"net/http"
)

// LogInfo  JSON
type logInfoRequest struct {
	ID      string `json:"ID"`
	Level   string `json:"level"`
	Message string `json:"message"`
}

type logInfoResponse struct {
	Code int    `json:"error"`
	Err  string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

// Logfile  JSON
type logfileRequest struct {
	ID    string `json:"ID"`
	Level string `json:"level"`
}

type logfileResponse struct {
	Filecontent string `json:"filecontent"`
	Err         string `json:"err,omitempty"`
}

// Client Id
type IDResponse struct {
	ID string `json:"ID"`
}

// FUNCTION      : decodeLogMessageRequest
// DESCRIPTION   : Translate a JSON object into an GO type
func decodeLogMessageRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request logInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// FUNCTION      : decodeLogfileRequest
// DESCRIPTION   : Translate a JSON object into an GO type
func decodeLogfileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request logfileRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// FUNCTION      : decodeLogMessageRequest
// DESCRIPTION   : Not necessary for the Generation of ID
func decodeIDGenerationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

// FUNCTION      : decodeLogMessageRequest
// DESCRIPTION   : Translate a Http response into a JSON oject
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
