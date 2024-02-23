package common

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// TODO:
// functions and err handling clean up
// should take filename as input?

// This script will take a user json file and then output it as a cairo file for the mapper
//var filename = "../../data/input.json"

func ConvertJsonToCairo(input string, dst string) {
	var data Result = read_json(input)
	var mat, vect = data.Matrix, data.Vector
	fmt.Printf("Matrix: %v\n", mat)
	fmt.Printf("Vector: %v\n", vect)
	var row, col = len(mat), len(mat[0])
	var vec_size = len(vect)
	fmt.Printf("row size: %v\n", row)
	fmt.Printf("col size: %v\n", col)
	fmt.Printf("vector size: %v\n", vec_size)
	assert(vec_size == col, "dimension mismatch")

	write_cairo_file(data, dst)
}

type Result struct {
	Matrix [][]int `json:"matrix"`
	Vector []int   `json:"vector"`
}

func assert(condition bool, message string) {
	if !condition {
		fmt.Fprintf(os.Stderr, "Assertion failed: %s\n", message)
		os.Exit(1)
	}
}
func read_json(filename string) Result {
	file, err := os.Open(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var data Result
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		// Handle the error, possibly log it
		log.Fatalf("Error decoding JSON: %v", err)
	}
	// Use the unmarshalled data

	return data
}

func write_cairo_file(data Result, filename string) {

	//opening the file to enter
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	//opening the data and matrix
	var matrix, vector = data.Matrix, data.Vector

	// Step 2: Write the header
	header := `use array::ArrayTrait;
use cairomap::matvectmult::{Matrix, Vec, matrixTrait, vecTrait, mapper, reducer, final_output};`
	file.WriteString(header)

	//Step 3: Write the matrix creation function

	// Write the matrix creation function
	file.WriteString("//return matrix, row, col\n")
	file.WriteString("fn create_matrix() -> (Matrix,u32,u32){\n")
	for i, row := range matrix {
		file.WriteString(fmt.Sprintf("    let row%d = array![", i+1))
		for j, val := range row {
			file.WriteString(fmt.Sprintf("%d", val))
			if j < len(row)-1 {
				file.WriteString(", ")
			}
		}
		file.WriteString("];\n")
	}

	file.WriteString("    let matrix_array = array![")
	for i := range matrix {
		file.WriteString(fmt.Sprintf("row%d", i+1))
		if i < len(matrix)-1 {
			file.WriteString(", ")
		}
	}
	// Dynamically generate the matrix dimensions
	file.WriteString("];\n")
	rows, cols := len(matrix), len(matrix[0])
	file.WriteString(fmt.Sprintf("    let mat = matrixTrait::init_array(%d, %d, @matrix_array);\n", rows, cols))
	file.WriteString(fmt.Sprintf("    (mat,%d,%d)\n", rows, cols))
	file.WriteString("}\n\n")

	// Write the vector creation function
	file.WriteString("//return row, vector_length\n")
	file.WriteString("fn create_vector()->(Vec,u32){\n")
	file.WriteString("    let vec_test = array![")
	for i, val := range vector {
		file.WriteString(fmt.Sprintf("%d", val))
		if i < len(vector)-1 {
			file.WriteString(", ")
		}
	}
	// Dynamically generate the vector length
	vectorLength := len(vector)
	file.WriteString("];\n")
	file.WriteString(fmt.Sprintf("    let vec = vecTrait::init_array(%d, @vec_test);\n", vectorLength))
	file.WriteString(fmt.Sprintf("    (vec,%d)\n", vectorLength))
	file.WriteString("}\n")

}
