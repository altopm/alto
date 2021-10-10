package utils

import (
	"strings"
)

func SplitString(str string, sep string) []string {
	strs := strings.Split(str, sep)
	return strs
}
