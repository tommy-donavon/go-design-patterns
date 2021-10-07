package models

import (
	"decorator/interfaces"
)

type ShifterReader struct {
	IO interfaces.IOReader
}

func (s *ShifterReader) ReadFile(fileName string) (string, error) {
	initial, err := s.IO.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	op := []rune{}
	for _, ch := range initial {
		op = append(op, ch-1)
	}
	return string(op), nil
}
