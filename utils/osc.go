package utils

import (
	"runtime"
)

func OsCheck() string {
	var os string = runtime.GOOS
	return os
}
