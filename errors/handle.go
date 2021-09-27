package errors

import (
	"fmt"

	. "github.com/logrusorgru/aurora"
)

func Handle(errMsg string) {
	fmt.Println(Red("\t Error! "), errMsg)
}
