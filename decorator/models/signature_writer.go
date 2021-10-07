package models

import (
	"decorator/interfaces"
	"fmt"
	"time"
)

type SignatureWriter struct {
	IO interfaces.IOWriter
}

func (s *SignatureWriter) WriteFile(input, fileName string) error {
	t := time.Now()
	return s.IO.WriteFile(fmt.Sprintf("%s \r\n\r\nTommy Donavon \r\n%d/%d/%d", input, t.Day(), t.Month(), t.Year()), fileName)
}
