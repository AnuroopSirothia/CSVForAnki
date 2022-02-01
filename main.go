package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Please give input and output file name.")
	}

	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	// Open input file which contains notes.
	notesFile := openNotesFile(inputFileName)

	// Parse file
	notesInCSVFormat := parseFile(notesFile)

	// Save file
	writeToCSVFile(notesInCSVFormat, outputFileName)

	// Print the file that is saved for debugging. This should enabled optionally.
	//	fmt.Print("Output File:\n", string(notesInCSVFormat))
}

func openNotesFile(fileToConvert string) *os.File {
	log.SetPrefix("csvforanki: ")
	log.SetFlags(0)

	var inputFile, e = os.Open(fileToConvert)
	if e != nil {
		log.Fatal(e)
	}

	return inputFile
}

func parseFile(file *os.File) []byte {
	scanner := bufio.NewScanner(file)
	var parsedData []byte

	for scanner.Scan() {
		line := scanner.Text()

		if line == "#" {
			continue
		}

		if line == "---" {
			parsedData = append(parsedData, parseQA(scanner)...)
		}
	} //for end
	return parsedData
}

func writeToCSVFile(data []byte, outputFileName string) {
	e := os.WriteFile(outputFileName+".csv", data, 0777)

	if e != nil {
		log.Fatal(e)
	}
}

// Parse the Question, Answer and Tag. Each tuple of question, answer and tag starts with --- and ends with ... like a YAML file.
func parseQA(scanner *bufio.Scanner) []byte {
	var currentLine, currentLineButOne, currentLineButTwo, questionLine string
	var answerLines strings.Builder

	var outputToSave []byte

	scanner.Scan()
	questionLine = scanner.Text()
	questionOutput := fmt.Sprintf("\n%q,", questionLine)
	outputToSave = append(outputToSave, []byte(questionOutput)...)

	for scanner.Scan() {
		currentLineButTwo = currentLineButOne
		currentLineButOne = currentLine
		currentLine = scanner.Text()

		if currentLine == "..." {
			answerLines.WriteString(currentLineButTwo)

			ansOutput := fmt.Sprintf("\"%s\",", answerLines.String())
			outputToSave = append(outputToSave, []byte(ansOutput)...)
			tagOutput := fmt.Sprintf("%q\n", currentLineButOne)
			outputToSave = append(outputToSave, []byte(tagOutput)...)
			break
		}

		ansOutput := fmt.Sprintf("%s\n", currentLineButTwo)
		answerLines.WriteString(ansOutput)

	} //for end

	return outputToSave
}
