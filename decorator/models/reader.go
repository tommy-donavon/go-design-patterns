package models

import "io/ioutil"

type FileReader struct{}

func (f *FileReader) ReadFile(fileName string) (string, error) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(fileBytes), nil
}
