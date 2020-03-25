/*
*  FILE          : service.go
*  PROJECT       : Assignment #3
*  PROGRAMMER    : Gabriel Gurgel
*  FIRST VERSION : 2020-03-24
*  DESCRIPTION   : Defines the service and all its methods
 */

package main

// Logopher provides logging capabilities
type Logopher interface {
	LogEvent(string, string, string) (int, error)
	RequestFile(string, string) (string, error)
	GenerateID() string
}
