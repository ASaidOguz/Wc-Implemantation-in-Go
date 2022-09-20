package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var fileName string

func main() {

	// First element in os.Args is always the program name,
	// So we need at least 2 arguments to have a file name argument.
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}
	// data is the file content, you can use it
	fmt.Println("File content is:")
	fmt.Println(string(data))
	fmt.Println("------------------------")

	fmt.Println("Line     Words    Characters")
	fmt.Println("----     -----    ----------")
	fileName = os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Err ", err)
	}
	scanner := bufio.NewScanner(file)
	lines, words, characters := 0, 0, 0
	for scanner.Scan() {
		lines++

		line := scanner.Text()
		characters += len(line)

		splitLines := strings.Split(line, " ")
		words += len(splitLines)
	}
	fmt.Printf("%2d  %8d  %8d      %s\n", lines, words, characters, fileName)
}
