package utils

import (
	"errors"
	"runtime"
)

func OsCheck() (string, error) {
	var os string = runtime.GOOS
	if os != "windows" && os != "linux" && os != "darwin" {
		var UnsupportedOSError error = errors.New("OS not supported!")
		return "", UnsupportedOSError
	}
	return os, nil
}
