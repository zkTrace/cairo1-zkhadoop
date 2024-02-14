package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// writeIntermediateFiles takes partitioned key-value pairs generated from a map task
// and writes them to intermediate files, one for each partition. The function returns
// a list of the generated file names. Each file is named using the map job number and partition index.
func writeIntermediateFiles(partitionedKva [][]KeyValue, mapJobNumber int) []string {
	intermediateFiles := make([]string, len(partitionedKva))
	for i, kva := range partitionedKva {
		// Construct a file name using the map job number and partition index.
		fileName := fmt.Sprintf("mr-%d-%d", mapJobNumber, i)
		intermediateFiles[i] = fileName

		// Create the file and check for errors.
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatalf("cannot create %v", fileName)
		}
		defer file.Close()

		// Encode the key-value pairs into the file using JSON format.
		jsonEncoder := json.NewEncoder(file)
		if err := jsonEncoder.Encode(kva); err != nil {
			log.Fatalf("cannot encode kva: %v", err)
		}
	}
	return intermediateFiles
}

// readIntermediateFiles takes a list of intermediate file names, reads them,
// and decodes the JSON content back into slices of KeyValue pairs. It aggregates
// all KeyValue pairs from all files into a single slice and returns it.
func readIntermediateFiles(files []string) []KeyValue {
	var intermediate []KeyValue
	for _, file := range files {
		// Read the file content.
		data, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("cannot read %v: %v", file, err)
		}

		// Decode the JSON content into a slice of KeyValue pairs.
		var kva []KeyValue
		if err := json.Unmarshal(data, &kva); err != nil {
			log.Fatalf("cannot unmarshal kva: %v", err)
		}

		// Append the decoded KeyValue pairs to the aggregate slice.
		intermediate = append(intermediate, kva...)
	}
	return intermediate
}

// writeReduceOutput takes the KeyValue pairs resulting from a reduce task,
// groups them by key, applies the reduce function, and writes the output
// to a final output file named using the reduce task number. Each line of the output
// file contains a key and its corresponding aggregated value.
func writeReduceOutput(intermediate []KeyValue, reduceNumber int, reducef func(string, []string) string) {
	// Construct the output file name using the reduce task number.
	outputFile := fmt.Sprintf("mr-out-%d", reduceNumber)
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("cannot create %v", outputFile)
	}
	defer file.Close()

	// Iterate over the sorted KeyValue pairs, aggregating values by key.
	i := 0
	for i < len(intermediate) {
		j := i + 1
		// Find all values for the current key.
		for j < len(intermediate) && intermediate[j].Key == intermediate[i].Key {
			j++
		}

		// Collect and reduce the values for the current key.
		values := make([]string, j-i)
		for k := i; k < j; k++ {
			values[k-i] = intermediate[k].Value
		}

		// Apply the reduce function and write the result to the output file.
		output := reducef(intermediate[i].Key, values)
		fmt.Fprintf(file, "%v %v\n", intermediate[i].Key, output)

		i = j // Move to the next group of KeyValue pairs.
	}
}
