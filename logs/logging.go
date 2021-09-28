package logs

import (
	"fmt"
	"os"
	"time"

	"github.com/altopm/alto/errors"
)

func CreateLogFile() {
	time := time.Now()
	logFile, err := os.Create("logs.log")
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
	file, err := os.Create("logs.log")
	if err != nil {
		errors.Handle(err.Error())
	}
	_, err = fmt.Fprintln(file, fmt.Sprintf("%s: %s", time.Format("2006-01-02 15:04:05.000000"), event))
	if err != nil {
		errors.Handle(err.Error())
	}
}
