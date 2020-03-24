package main

import (
	"errors"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// From config file
const timeLocal = "UTC"
const logPath = "./logfiles/" // DEBUG: Add the id here
const ext = ".log"
const temporaryID = 123 // DEBUG:

// Service struct
type logopher struct{}

// Service Methods
func (logopher) LogEvent(level string, message string) (string, error) {
	// Check for empty log message
	if message == "" {
		return "", errors.New("Empty Log Message")
	}
	// Remove invalid characters
	message = compileLogMessage(message)

	// Valid log level
	level = strings.ToLower(level)
	validLevels := getValidLevels()
	_, found := validLevels[level]
	if !found {
		return "", errors.New("Invalid log level")
	}

	// Write to file
	appendToLog(level, message)
	return message, nil
}

// Append to the file
func appendToLog(level string, logMessage string) {
	levelPath := logPath + level + ext
	allPath := logPath + strconv.Itoa(temporaryID) + ext

	logFile, err := os.OpenFile(levelPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	allLogs, err := os.OpenFile(allPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	location, _ := time.LoadLocation(timeLocal)
	now := time.Now().In(location).String()
	logFile.WriteString("[" + now + "] [" + strings.ToUpper(level) + "] " + logMessage + "\n")
	allLogs.WriteString("[" + now + "] [" + strings.ToUpper(level) + "] " + logMessage + "\n")
	logFile.Close()
	allLogs.Close()
}

func getValidLevels() map[string]struct{} {
	l := make(map[string]struct{})
	l["debug"] = struct{}{}
	l["warn"] = struct{}{}
	l["error"] = struct{}{}
	l["fatal"] = struct{}{}
	return l
}

func compileLogMessage(message string) string {
	space := regexp.MustCompile(`\s+`)
	reg, err := regexp.Compile("[^a-zA-Z0-9._ ]+")
	if err != nil {
		log.Fatal(err)
	}

	message = reg.ReplaceAllString(message, "")
	message = space.ReplaceAllString(message, " ")
	return message
}
