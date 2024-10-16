package logger

import (
	"fmt"
	"os"
)

type LOGLEVEL int // Log levels
type Path string  // Path of dir
type LogFileVariablesFormat struct {
	dir                    Path
	filename               Path
	isFileLoggingAvailable bool
	logType                LOGLEVEL
}

var LogFileVariables LogFileVariablesFormat

const (
	LOG LOGLEVEL = 1 + iota
	FILE
	MIXED
)

const LOG_FOLDER_NAME string = "logs"

func SetUpLog(dirPath Path, fileName Path, logType LOGLEVEL) {

	// create log directory
	os.Mkdir(LOG_FOLDER_NAME, 0777)

	LogFileVariables.dir = dirPath + "/" + Path(LOG_FOLDER_NAME) + "/"
	LogFileVariables.filename = fileName
	LogFileVariables.isFileLoggingAvailable = true
	LogFileVariables.logType = logType

	// create log file
	path := fmt.Sprintf("%s%s.txt", LogFileVariables.dir, LogFileVariables.filename)
	f, e := os.Create(path)

	if e == nil {
		Log("Logger is configured successfully!")
	}

	defer f.Close()
}

func Log(content string) {

	switch LogFileVariables.logType {
	case LOG:
		logToConsole(content)

	case FILE:
		defer logToFile(content)

	case MIXED:
		logToConsole(content)
		defer logToFile(content)
	}
}

func logToConsole(content string) {
	fmt.Printf("%s\n", content)
}

func logToFile(content string) error {

	var isDirPresent bool = len(LogFileVariables.dir) > 0
	var isFileNamePresent bool = len(LogFileVariables.filename) > 0

	if !(isDirPresent && isFileNamePresent && LogFileVariables.isFileLoggingAvailable) {

		LogFileVariables.isFileLoggingAvailable = false

		logToConsole("Logger is not configured for logging on to file, Switching back to console logging")

		return nil
	} else {
		path := fmt.Sprintf("%s%s.txt", LogFileVariables.dir, LogFileVariables.filename)

		f, _ := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		_, e := f.Write([]byte(content + "\n"))

		defer f.Close()
		return e

	}
}
