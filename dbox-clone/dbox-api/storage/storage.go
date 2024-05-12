package storage

import (
	"encoding/base64"
	"os"
)

func MakeDirectory(name string) bool {
	var status bool
	err := os.Mkdir(os.Getenv("STORE")+"/"+name, 0777)
	if err != nil {
		status = false
	}
	status = true

	return status
}
func CreateFile(directory string, filename string, fileType string, data string) (string, error) {
	var path string

	decoder, err := base64.StdEncoding.DecodeString(data)

	if err != nil {
		return "", err
	}
	path = directory + "/" + filename + "." + fileType
	f, err := os.Create(path)

	if err != nil {
		return "", err
	}
	defer f.Close()

	if _, err := f.Write(decoder); err != nil {
		return "", err
	}

	return path, nil
}
