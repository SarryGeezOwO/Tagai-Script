package main

import (
	"fmt"
	"os"
	"bufio"
	"path/filepath"
	syn "tagai-script/syntax"
	er "tagai-script/error"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: <file-name>.tgs")
		return
	}

	sourceFile := args[1]
	if !validateExtension(sourceFile) {
		fmt.Println("Invalid file format:", sourceFile, "expected .tgs file")
		return
	}

	RunFile(sourceFile)
}



// Returns true if the given File ends with .tgs extension
func validateExtension(fileName string) bool {
	extension := filepath.Ext(fileName)
	return extension == ".tgs"
}



// Reads a given .tgs file and invokes the Run() function
func RunFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(69)
	}
	defer file.Close()

	// Read file line by line and create
	// The source code that will be given to Run()
	scanner := bufio.NewScanner(file)
	sourceStr := "" 

	for scanner.Scan() {
		sourceStr += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	Run(sourceStr)
}



func Run(source string) {
	tokens := syn.Tokenize(source)

	for _, token := range tokens {
		fmt.Println("Line:", token.Line, "Type:", token.Type, token.Lexeme, token.Literal)
	}

	if er.ErrorPresent {
		fmt.Println("Error occured, evaluation stopped.")
	}
}