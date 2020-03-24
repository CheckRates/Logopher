package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type logInfoRequest struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

type logInfoResponse struct {
	ErrorMessage string `json:"error"`
	Err          string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

func decodeLogMessageRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request logInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
