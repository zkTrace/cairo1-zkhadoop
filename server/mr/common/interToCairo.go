package common

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// TODO:
// functions and err handling clean up
// should take filename as input

// This script will take the intermediary file and then write it for the reducer
func ConvertIntermediateToCairo() {
	var read_fileName = "../../files/map_files/mapper_res.txt"
	read_intermediary(read_fileName)

}

// json struct to read from
type Data struct {
	IntermediaryValues [][]int `json:"intermediary_values"`
}

// This function will read the intermediary file name
func read_intermediary(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var jsonString strings.Builder
	inJSON := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Check if this line marks the beginning of the JSON content
		if strings.Contains(line, "{") {
			inJSON = true
		}

		if inJSON {
			// Add the line to the JSON string
			jsonString.WriteString(line)
		}

		// Check if this line marks the end of the JSON content
		if strings.Contains(line, "}") {
			break // Stop reading the file once the end of the JSON content is reached
		}
	}

	// Remove the trailing semicolon if present
	jsonStringStr := strings.TrimSuffix(jsonString.String(), ";")

	// Unmarshal the JSON string into a Go struct
	var data Data
	if err := json.Unmarshal([]byte(jsonStringStr), &data); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Print the result to verify
	fmt.Printf("Intermediary Values: %+v\n", data.IntermediaryValues)
}
