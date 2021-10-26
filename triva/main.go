package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"triva/models"
	"triva/utils"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	quiz    = models.NewSection("Main Quiz")
)

func main() {
	runProgram := true
	for runProgram {
		userInput := utils.MakeMenu("Make Quiz", "Take Quiz")
		switch userInput {
		case 1:
			quizChoice := utils.MakeMenu("Make New Section", "Make New Question")
			switch quizChoice {
			case 1:
				sectionChoice := utils.MakeMenu("Make New Main Section", "Make New Sub Section")
				switch sectionChoice {
				case 1:
					if scanner.Scan() {
						quiz.Add(models.NewSection(scanner.Text()))
					}
				case 2:
					fmt.Println("Choice a Section to add a sub section to: ")
					st, _ := quiz.GetIterator()
					tempCol := []string{}
					for st.HasNext() {
						val, ok := st.GetNext().(*models.Section)
						if ok {
							tempCol = append(tempCol, val.GetName())
						}
					}
					subSelection := utils.MakeMenu(tempCol...)
					st, _ = quiz.GetIterator()
					for st.HasNext() {
						val, ok := st.GetNext().(*models.Section)
						if ok && val.GetName() == tempCol[subSelection-1] {
							fmt.Println("Enter the name of your sub section")
							if scanner.Scan() {
								val.Add(models.NewSection(scanner.Text()))
							}
						}
					}
				}
			case 2:
				fmt.Println("Choose a Section to add a question to: ")
				st, _ := quiz.GetIterator()
				tempCol := []string{}
				for st.HasNext() {
					val, ok := st.GetNext().(*models.Section)
					if ok {
						tempCol = append(tempCol, val.GetName())
					}
				}
				subSelection := utils.MakeMenu(tempCol...)
				st, _ = quiz.GetIterator()
				for st.HasNext() {
					val, ok := st.GetNext().(*models.Section)
					if ok && val.GetName() == tempCol[subSelection-1] {
						ques, corr, poss := "", "", ""

						fmt.Print("Enter your question: ")
						if scanner.Scan() {
							ques = scanner.Text()
						}
						fmt.Print("Enter the correct answer: ")
						if scanner.Scan() {
							corr = scanner.Text()

						}
						fmt.Print("Enter in other possible answers(seperated by a , )")
						if scanner.Scan() {
							poss = scanner.Text()
						}
						possAsList := strings.Split(poss, ",")
						val.Add(models.NewQuestion(ques, corr, possAsList...))
					}
				}

			case 0:
				userInput = 0
			}
		case 2:
			if quiz.Size() == 0 {
				quiz = buildSampleTest()
			}
			quiz.AskQuestion()
			fmt.Printf("Score: %d out of %d\n", models.Score, models.OutOf)
		case 0:
			runProgram = false
		}
	}

}

func buildSampleTest() *models.Section {
	/// main test
	tr := models.NewSection("Main Quiz")

	// Sections
	sci := models.NewSection("Science")

	// SubSections
	comp_sci := models.NewSection("Computer Science")
	math_sci := models.NewSection("Mathematic Science")
	gen_sci := models.NewSection("General Science")

	//Questions
	comp_sci.Add(models.NewQuestion("The Harvard architecture for micro-controllers added which additional bus?", "Instruction", "Address", "Data", "Control"))
	comp_sci.Add(models.NewQuestion("What amount of bits commonly equals one byte", "8", "1", "2", "64"))
	comp_sci.Add(models.NewQuestion("What was the name given to Android 4.3", "Jelly Bean", "Lollipop", "Nutella", "Froyo"))

	math_sci.Add(models.NewQuestion("What is the area of a circle with a diameter of 20 inches if pi = 3.1415?", "341.15 Inches", "3141.5 Inches", "380.1215 Inches"))
	math_sci.Add(models.NewQuestion("What is the square root of 49", "7", "4", "12", "9"))
	math_sci.Add(models.NewQuestion("The decimal number 32 in hexadecimal would be what?", "1F", "3D", "2E", "1B"))

	gen_sci.Add(models.NewQuestion("Which gas forms about 78% of the Earth's atmosphere?", "Nitrogen", "Oxygen", "Argon", "Carbon Dioxide"))
	gen_sci.Add(models.NewQuestion("Au on the Periodic Table refers to which element?", "Gold", "Silver", "Oxygen", "Nickel"))
	gen_sci.Add(models.NewQuestion("What is the atomic number of the element Strontium?", "38", "73", "47", "11"))

	sci.Add(comp_sci)
	sci.Add(math_sci)
	sci.Add(gen_sci)
	tr.Add(sci)

	return tr

}
