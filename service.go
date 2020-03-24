package main

// Logopher provides logging capabilities
type Logopher interface {
	LogEvent(string, string) (string, error)
}
