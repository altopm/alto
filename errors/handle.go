package errors

import (
	"fmt"
	"os"

	. "github.com/logrusorgru/aurora"
)

func Handle(errMsg string) {
	fmt.Println(Red("\t Error! "), errMsg)
	os.Exit(1)
}
