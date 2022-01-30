package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
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

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "#" {
			continue
		}

		if line == "---" {
			formatQA(scanner)
		}
	}
}

// Format the Question, Answer and Tag
func formatQA(scanner *bufio.Scanner) {
	var currentLine, currentLineButOne, currentLineButTwo string
	var questionLine string
	var answer strings.Builder

	scanner.Scan()
	questionLine = scanner.Text()
	fmt.Printf("\n%q,", questionLine)

	for scanner.Scan() {
		currentLineButTwo = currentLineButOne
		currentLineButOne = currentLine
		currentLine = scanner.Text()

		if currentLine == "..." {
			answer.WriteString(currentLineButTwo)
			fmt.Printf("%q,", answer.String())
			fmt.Printf("%q,\n", currentLineButOne)
			break
		}

		//		fmt.Println("Writing to answer buffer ", currentLineButTwo)
		answer.WriteString(currentLineButTwo)

	} //for end
}
