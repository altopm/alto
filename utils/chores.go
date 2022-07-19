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

func GetRegistryList() ([]string, error) {
	list := []byte{}
	if !DoesDirectoryExist("/var/alto/registry/") {
		err := os.MkdirAll("/var/alto/registry", 0755)
		if err != nil {
			return nil, err
		}

		err = os.WriteFile("/var/alto/registry/entries.json", []byte("[]"), 0644)

		if err != nil {
			return nil, err
		}

		return []string{string(list)}, nil
	}

	file, err := os.Open("/var/alto/registry/entries.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fileData, err := file.Read(make([]byte, 50))
	if err != nil {
		return nil, err
	}
	return []string{string(rune(fileData))}, nil
}
