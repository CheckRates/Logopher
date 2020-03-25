/*
*  FILE          : logic.go
*  PROJECT       : Assignment #3
*  PROGRAMMER    : Gabriel Gurgel
*  FIRST VERSION : 2020-03-24
*  DESCRIPTION   : Defines the logic behind each function used for the endpoints
 */

package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/rs/xid"
)

// From config file
const timeLocal = "UTC"
const logPath = "./logfiles/" // DEBUG: Add the id here
const ext = "log"

// Service struct
type logopher struct{}

func (logopher) GenerateID() string {
	generatedID := xid.New().String()
	os.MkdirAll(logPath+generatedID, os.ModePerm)

	return generatedID
}

// FUNCTION      : RequestFile
// DESCRIPTION   : Retrieves the log file by level
//
// PARAMETERS    :
//		String id: Id of the client that is logging
//		string level: log level of the message
//
// RETURNS       :
//	Returns the contents od the logfile if sucecsss, error description if it doesnt
func (logopher) RequestFile(id string, level string) (string, error) {
	level = strings.ToLower(level)

	if level == "all" {
		level = id
	} else {
		// Validate log level
		validLevels := getValidLevels()
		_, found := validLevels[level]
		if !found {
			return "", errors.New("Invalid log level")
		}
	}

	// Open file and put all the contents into a byte array
	requestedPath := logPath + id + "/" + level + "." + ext
	contents, err := ioutil.ReadFile(requestedPath)
	if err != nil {
		return "404", errors.New("File not found")
	}

	return string(contents), nil
}

// FUNCTION      : LogEvent
// DESCRIPTION   : Logs an event with its level
//
// PARAMETERS    :
//		String id: Id of the client that is logging
//		string level: log level of the message
//		string message: log message itself
//
// RETURNS       :
//	0 if sucess, an negative error and description if not
func (logopher) LogEvent(id string, level string, message string) (int, error) {
	// Check if the user exists
	if !checkIfClientExists(id) {
		return -1, errors.New("Client does not exist -- Please generate your key")
	}

	// Check for empty log message
	if message == "" {
		return -2, errors.New("Empty Log Message")
	}
	// Remove invalid characters
	message = compileLogMessage(message)

	// Valid log level
	level = strings.ToLower(level)
	validLevels := getValidLevels()
	_, found := validLevels[level]
	if !found {
		return -3, errors.New("Invalid log level")
	}

	// Write to file
	appendToLog(id, level, message)
	return 1, nil
}

// FUNCTION      : appendToLog
// DESCRIPTION   : Handles the file io of the LogEvent
//
// PARAMETERS    :
//		String id: Id of the client that is logging
//		string level: log level of the message
//		string message: log message itself
//
// RETURNS       :
// Nada
func appendToLog(id string, level string, logMessage string) {
	os.MkdirAll(logPath+id, os.ModePerm)
	levelPath := logPath + id + "/" + level + "." + ext
	allPath := logPath + id + "/" + id + "." + ext

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

// FUNCTION      : getValidLevels
// DESCRIPTION   : Contains all the valid log levels
//
// PARAMETERS    :
// 	Nothing
//
// RETURNS       :
// A map containing all valid log levels
func getValidLevels() map[string]struct{} {
	l := make(map[string]struct{})
	l["debug"] = struct{}{}
	l["warn"] = struct{}{}
	l["error"] = struct{}{}
	l["fatal"] = struct{}{}
	return l
}

// FUNCTION      : compileLogMessage
// DESCRIPTION   : Contains all the valid log levels
//
// PARAMETERS    :
// 	Nothing
//
// RETURNS       :
// A map containing all valid log levels
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

// FUNCTION      : checkIfClientExists
// DESCRIPTION   : Check if the client exists
//
// PARAMETERS    :
// 	Nothing
//
// RETURNS       :
// True if the client exists, false otherwise
func checkIfClientExists(id string) bool {
	clients, err := ioutil.ReadDir(logPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, clientID := range clients {
		if id == clientID.Name() {
			return true
		}
	}
	return false
}
