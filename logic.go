package main

import (
	"errors"
)

// Service struct
type logopher struct{}

// Service Methods
func (logopher) LogEvent(message string) (string, error) {
	if message == "" {
		return "", errors.New("Empty Log Message")
	}
	return message, nil
}
