package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Open input file which contains notes.
	notesFile := openNotesFile()

	// Parse file
	notesInCSVFormat := parseFile(notesFile)

	// Save file
	writeToCSVFile(notesInCSVFormat)
}

func openNotesFile() *os.File {
	log.SetPrefix("csvforanki: ")
	log.SetFlags(0)

	if len(os.Args) != 2 {
		log.Fatal("Please give file name.")
	}

	fileToConvert := os.Args[1]

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

func writeToCSVFile(data []byte) {
	e := os.WriteFile("output/import_ready.txt", data, 0777)

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
	fmt.Print(questionOutput)
	outputToSave = append(outputToSave, []byte(questionOutput)...)

	for scanner.Scan() {
		currentLineButTwo = currentLineButOne
		currentLineButOne = currentLine
		currentLine = scanner.Text()

		if currentLine == "..." {
			answerLines.WriteString(currentLineButTwo)

			ansOutput := fmt.Sprintf("%q,", answerLines.String())
			fmt.Print(ansOutput)
			outputToSave = append(outputToSave, []byte(ansOutput)...)
			tagOutput := fmt.Sprintf("%q,\n", currentLineButOne)
			fmt.Print(tagOutput)
			outputToSave = append(outputToSave, []byte(tagOutput)...)
			return outputToSave
		}

		//		fmt.Println("Writing to answer buffer ", currentLineButTwo)
		answerLines.WriteString(currentLineButTwo)

	} //for end

	return nil
}
