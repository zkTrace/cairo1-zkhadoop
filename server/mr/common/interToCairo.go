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
func ConvertIntermediateToCairo(input string, dst string) {
	// var read_fileName = "../../files/map_files/mapper_res.txt"
	var data Data = read_intermediary(input)
	// write_intermediary(data, "matvecdata_reducer.cairo")
	write_intermediary(data, dst)

}

// json struct to read from
type Data struct {
	IntermediaryValues [][]int `json:"intermediary_values"`
}

// This function will read the intermediary file name
// This function will read the intermediary file name
func read_intermediary(filename string) Data {
	file, err := os.Open(filename)
	var data Data

	if err != nil {
		fmt.Println("Error opening file:", err)
		return data
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
	if err := json.Unmarshal([]byte(jsonStringStr), &data); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return data

	}

	// Print the result to verify
	fmt.Printf("Intermediary Values: %+v\n", data.IntermediaryValues)
	return data
}

func write_intermediary(data Data, dst string) {
	file, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the static part of your desired content
	file.WriteString("pub mod inter_val{\n")
	file.WriteString("    pub fn inter_result()->Array<(u32, felt252)>{\n")
	file.WriteString("        let inter_res = array![")

	// Dynamically write the data
	for i, pair := range data.IntermediaryValues {
		if i > 0 {
			file.WriteString(", ")
		}
		file.WriteString(fmt.Sprintf("(%d,%d)", pair[0], pair[1]))
	}

	// Write the closing part of your desired content
	file.WriteString("];\n")
	file.WriteString("        inter_res\n")
	file.WriteString("    }\n")
	file.WriteString("}\n")

	if err != nil {
		panic(err)
	}

	fmt.Println("Data written to matvecdata_reducer.cairo successfully.")
}
