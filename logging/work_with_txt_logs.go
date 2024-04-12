package logging

import (
	"log"
	"os"
)

var (
	WarningLog *log.Logger
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
)

func Init() {
	logFile, err := os.OpenFile("appLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLog = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLog = log.New(logFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Execute_txt_log() {
	Init()
	InfoLog.Println("Starting the application...")
	InfoLog.Println("Something has occurred...")
	WarningLog.Println("WARNING!!!")
	ErrorLog.Println("Some error has occurred...")
}
