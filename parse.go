package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	print("Pipe the output of this program to jq -R \"fromjson?\"")
	parser(os.Stdin)
}

func parser(r io.Reader) error {

	scanner := bufio.NewScanner(bufio.NewReader(r))

	for scanner.Scan() {
		// Regex set
		re := regexp.MustCompile(`{.*?}`)

		// Match the string
		toCheck := re.FindAll([]byte(scanner.Text()), -1)

		// Check if the byte array is empty
		if toCheck != nil {

			// A byte array
			bytesToProcess := toCheck[0]

			// Convert it to strings
			stringToProcess := string(bytesToProcess)

			// Print to Stdout
			fmt.Println(stringToProcess)

		}

	}
	return nil
}
