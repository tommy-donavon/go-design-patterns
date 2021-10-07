package main

import (
	"decorator/models"
	"fmt"
)

func main() {

	//FILE WRITING
	/*
		CASE ONE: Shifter decorator wraps Signature decorator

		RESULT: the Shifter decorator will call it's WriteFile method first
				shifting all chars up by one and then call the Signature decorator
				to append my name and todays date to that value
	*/
	base_writer := &models.FileWriter{}
	writer_sig := &models.SignatureWriter{
		IO: base_writer,
	}
	writer_sig_shift := &models.ShifterWriter{
		IO: writer_sig,
	}
	writer_sig_shift.WriteFile("ab", "output.txt")

	/*
		CASE TWO: Signature decorator wraps Shifter decorator

		RESULT: The Signature decorator will append my name and todays
		date to the initial value before calling the Shifter decorator's
		WriteFile method to shift all chars meaning that the signature
		is also altered by the Shifter decorator
	*/

	base_writer2 := &models.FileWriter{}
	writer_shift := &models.ShifterWriter{
		IO: base_writer2,
	}
	writer_shift_sig := &models.SignatureWriter{
		IO: writer_shift,
	}
	writer_shift_sig.WriteFile("ab", "output2.txt")

	//FILE READING
	base_reader := &models.FileReader{}
	reader_shift := &models.ShifterReader{
		IO: base_reader,
	}
	value, _ := reader_shift.ReadFile("output2.txt")
	fmt.Println(value)

	endings_reader := &models.LineEndingsReader{
		IO: base_reader,
	}

	eValue, _ := endings_reader.ReadFile("output.txt")
	base_writer.WriteFile(eValue, "output3.txt")

}
