package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseMapLine(line string) map[string]string {
	// Remove outer curly braces from the line
	line = line[1 : len(line)-1]

	// Split the line into key-value pairs
	pairs := strings.Split(line, ",")

	m := make(map[string]string)

	for _, pair := range pairs {
		// Split each pair into key and value

		kv := strings.SplitN(pair, ":", 2)

		// Remove leading/trailing spaces from the key
		key := strings.TrimSpace(kv[0])

		// Remove leading/trailing spaces from the value if it is not " "
		var value string
		if len(kv) > 1 {

			value = strings.TrimSpace(kv[1])
			if value == "" {
				value = " "
			}
		} else {
			value = " "
		}

		// Add the key-value pair to the map
		m[key] = value
	}

	return m
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	// Read the first line as a map
	scanner.Scan()
	mapLine := scanner.Text()
	codingTable := parseMapLine(mapLine)

	// Read the second line as a string
	scanner.Scan()
	encodedText := scanner.Text()

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	decodedText := decodeHuffman(encodedText, codingTable)
	fmt.Println(decodedText) // Output: ABEDABEC
}
