package ui

import (
	"fmt"

	. "github.com/logrusorgru/aurora"
)

func SuccessText(scMsg string) {
	fmt.Println(Green("\tSuccess! "), scMsg)
}
