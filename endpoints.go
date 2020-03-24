package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeLogInfoEndpoint(svc Logopher) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(logInfoRequest)
		errorMessage, err := svc.LogEvent(req.Level, req.Message)
		if err != nil {
			return logInfoResponse{errorMessage, err.Error()}, nil
		}
		return logInfoResponse{errorMessage, ""}, nil
	}
}

//RequestLogs(string)
