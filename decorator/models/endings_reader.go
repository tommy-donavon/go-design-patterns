package models

import (
	"decorator/interfaces"
	"strings"
)

type LineEndingsReader struct {
	IO interfaces.IOReader
}

func (s *LineEndingsReader) ReadFile(fileName string) (string, error) {
	initial, err := s.IO.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	initial = strings.ReplaceAll(initial, "\r\n", "\n")
	return initial, nil

}
