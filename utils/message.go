package utils

import (
	"fmt"

	aurora "github.com/logrusorgru/aurora"
)

func MessageNeutral(message string) {
	fmt.Println(aurora.Cyan(message))
}
func MessageWarning(message string) {
	fmt.Println(aurora.Yellow(message))
}
func MessageSuccess(scMsg string) {
	fmt.Println(aurora.Green("\tSuccess! "), scMsg)
}
func HackerText(message string) {
	fmt.Println(aurora.Green(message))
}