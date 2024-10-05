package utils

import (
    "log"
    "os"
)

var (
    InfoLogger  *log.Logger
    ErrorLogger *log.Logger
)

func init() {
    file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalln("Failed to open log file:", err)
    }

    InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogInfo(message string) {
    InfoLogger.Println(message)
}

func LogError(message string) {
    ErrorLogger.Println(message)
}
