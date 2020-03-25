/*
*  FILE          : endpoints.go
*  PROJECT       : Assignment #3
*  PROGRAMMER    : Gabriel Gurgel
*  FIRST VERSION : 2020-03-24
*  DESCRIPTION   : Defines endpoints for each request, each assigning to the funtion
*				  peformed in that endpoint
 */
package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// FUNCTION      : makeLogInfoEndpoint
// DESCRIPTION   : Makes the endpoint for logging a message
//
// PARAMETERS    :
//		Logopher svc : Service structure
//
// RETURNS       :
//	The access point to the functionality (endpoint)
func makeLogInfoEndpoint(svc Logopher) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(logInfoRequest)
		retCode, err := svc.LogEvent(req.ID, req.Level, req.Message)
		if err != nil {
			return logInfoResponse{retCode, err.Error()}, nil
		}
		return logInfoResponse{retCode, ""}, nil
	}
}

// FUNCTION      : makeLogfileRequestEndpoint
// DESCRIPTION   : Makes the endpoint for requesting the logfiles
//
// PARAMETERS    :
//		Logopher svc : Service structure
//
// RETURNS       :
//	The access point to the functionality (endpoint)
func makeLogfileRequestEndpoint(svc Logopher) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(logfileRequest)
		contents, err := svc.RequestFile(req.ID, req.Level)
		if err != nil {
			return logfileResponse{contents, err.Error()}, nil
		}

		return logfileResponse{contents, ""}, nil
	}
}

// FUNCTION      : makeGenerateIDEndpoint
// DESCRIPTION   : Makes the endpoint for generating the client id
//
// PARAMETERS    :
//		Logopher svc : Service structure
//
// RETURNS       :
//	The access point to the functionality (endpoint)
func makeGenerateIDEndpoint(svc Logopher) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		id := svc.GenerateID()
		return IDResponse{id}, nil
	}
}
