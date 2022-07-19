package logs

import (
	"fmt"
	"os"
	"time"

	"github.com/altopm/alto/errors"
)

func CreateLogFile() {
	time := time.Now()
	logFile, err := os.Create("alto-install.log")
	if err != nil {
		errors.Handle(err.Error())
	}
	_, err = fmt.Fprintf(logFile, fmt.Sprintf("%s: Created logfile\n", time.Format("2006-01-02 15:04:05.000000")))
	if err != nil {
		errors.Handle(err.Error())
	}
}
func AppendLog(event string) {
	time := time.Now()

	log, err := os.OpenFile("alto-install.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		errors.Handle(err.Error())
	}
	defer log.Close()

	_, err = log.WriteString(fmt.Sprintf("%s: %s\n", time.Format("2006-01-02 15:04:05.000000"), event))
	if err != nil {
		errors.Handle(err.Error())
	}
}
