package utils_test

import (
	"regexp"
	"runtime"
	"testing"

	"github.com/altopm/alto/utils"
)

func TestHelloName(t *testing.T) {
	os := runtime.GOOS
	want := regexp.MustCompile(`\b` + os + `\b`)
	resp, err := utils.OsCheck()
	if !want.MatchString(resp) || err != nil {
		t.Fatalf(`utils.OsCheck = %q, %v, want match for %#q, nil`, resp, err, want)
	}
}
