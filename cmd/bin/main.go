package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/iamvineettiwari/go-diff/tools"
)

func readFile(filename string) []byte {
	content, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	return content
}

func splitLines(content []byte) []string {
	return strings.Split(string(content), "\n")
}

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		log.Fatal("Invalid arguments")
	}

	firstContent := splitLines(readFile(args[0]))
	secondContent := splitLines(readFile(args[1]))

	diff := tools.FindDiff(firstContent, secondContent)

	for _, item := range diff {
		fmt.Println(item)
	}
}
