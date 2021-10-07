package models

import "os"

type FileWriter struct{}

func (f *FileWriter) WriteFile(input, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(input)
	return err
}
