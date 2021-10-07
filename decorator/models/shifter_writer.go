package models

import (
	"decorator/interfaces"
)

type ShifterWriter struct {
	IO interfaces.IOWriter
}

func (s *ShifterWriter) WriteFile(input, filename string) error {
	op := []rune{}
	for _, ch := range input {
		op = append(op, ch+1)
	}
	return s.IO.WriteFile(string(op), filename)
}
