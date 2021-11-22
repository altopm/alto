package utils

import (
	"os"
)

func DoesDirectoryExist(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
