package utils

import (
	"log"
	"os"
)

/* var (
	logfile, _ = os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logger     = log.New(logfile, "", log.LstdFlags)
)

func init() {
	log.SetOutput(logfile)
}

func LogMessage(message string) {
	logger.Println(message)
}
*/

var (
	logfile, _ = os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logger     = log.New(logfile, "", log.LstdFlags)
)

func init() {
	log.SetOutput(logfile)
}

func LogMessage(message string) {
	logger.Println(message)
}
