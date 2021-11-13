package utils

import (
	"fmt"

	"github.com/cavaliercoder/grab"
)

func GetPackage(packageName string) (*grab.Response, error) {
	resp, err := grab.Get(fmt.Sprintf("/var/alto/installs/%s", packageName), fmt.Sprintf("https://flags.altopkg.com/%s/%s.tar.gz", packageName, packageName))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
