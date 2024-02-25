package common

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

/**
	function: aggreate
		- input: string
		- output: write to aggDst
**/
func AggregateMapperCairo(aggDst string) {
	fmt.Println("====== Begin Combining all MAPPER Cairo Functions ======", )

	header := `
use array::ArrayTrait;
use core::dict::Felt252DictTrait;
use core::traits::TryInto;
use core::traits::Into;
use option::OptionTrait;
use core::debug::PrintTrait;
use core::fmt::Formatter;
`

	// 3 files lib.cairo, matvectdata_mapper.cairo, matvectmult.cairo
	mult_cairo, mult_cairo_error := read_cairo("/app/cairo/map/src/matvectmult.cairo")	
	data_cairo, data_cairo_error := read_cairo("/app/cairo/map/src/matvecdata_mapper.cairo")
	lib_cairo, lib_cairo_error := read_cairo("/app/cairo/map/src/lib.cairo")

	if mult_cairo_error != nil || data_cairo_error != nil || lib_cairo_error != nil {
		fmt.Println("Error reading file")
		return
	}
	
	file, err := os.Create(aggDst)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	file.WriteString(header)
	file.WriteString(removeHeaders(mult_cairo))
	file.WriteString(removeHeaders(data_cairo))
	file.WriteString(removeHeaders(lib_cairo))
}

func AggregateReducerCairo(aggDst string) {
	fmt.Println("====== Begin Combining all REDUCER Cairo Functions ======", )

	header := `
use array::ArrayTrait;
use core::dict::Felt252DictTrait;
use core::traits::TryInto;
use core::traits::Into;
use option::OptionTrait;
use core::debug::PrintTrait;
use core::fmt::Formatter;
`

	// 3 files lib.cairo, matvectdata_mapper.cairo, matvectmult.cairo
	mult_cairo, mult_cairo_error := read_cairo("/app/cairo/reducer/src/matvectmult.cairo")	
	data_cairo, data_cairo_error := read_cairo("/app/cairo/reducer/src/matvecdata_reducer.cairo")
	lib_cairo, lib_cairo_error := read_cairo("/app/cairo/reducer/src/lib.cairo")

	if mult_cairo_error != nil || data_cairo_error != nil || lib_cairo_error != nil {
		fmt.Println("Error reading file")
		return
	}
	
	file, err := os.Create(aggDst)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	file.WriteString(header)
	file.WriteString(removeHeaders(mult_cairo))
	file.WriteString(data_cairo)
	file.WriteString(removeHeaders(lib_cairo))
}

func read_cairo(filename string) (string, error) {
	file, err := os.Open(filename)
    if err != nil {
        // Handle the error
        fmt.Println("Error opening file:", err)
        return "failed", err
    }
    defer file.Close() // Make sure to close the file when done

    // Read the file content into a string
    // The io.ReadAll function reads the whole file into memory
    data, err := io.ReadAll(file)
    if err != nil {
        // Handle the error
        fmt.Println("Error reading file:", err)
        return "failed", err
    }

    // Convert data to a string
    content := string(data)

		return content, nil
}

func removeHeaders(input string) string {
	startDelimiter := `// === Header Start ===`
	endDelimiter := `// === Header End ===`


	// Find the start index of the delimiter
    startIdx := strings.Index(input, startDelimiter)
    if startIdx == -1 {
        // Start delimiter not found; return the original string
        return input
    }

    // Adjust startIdx to include the delimiter itself in the removal
    startIdx += len(startDelimiter)

    // Find the end index of the delimiter
    endIdx := strings.Index(input[startIdx:], endDelimiter)
    if endIdx == -1 {
        // End delimiter not found; return the original string
        return input
    }

    // Adjust endIdx to refer to the position in the original string
    // and to include the delimiter itself in the removal
    endIdx += startIdx + len(endDelimiter)

    // Concatenate the part before the start delimiter and the part after the end delimiter
    // return input[:startIdx] + input[endIdx:]
		return input[endIdx:]
}

// func ConvertJsonToCairo(input string, dst string) {
// 	var data Result = read_json(input)
// 	var mat, vect = data.Matrix, data.Vector
// 	fmt.Printf("Matrix: %v\n", mat)
// 	fmt.Printf("Vector: %v\n", vect)
// 	var row, col = len(mat), len(mat[0])
// 	var vec_size = len(vect)
// 	fmt.Printf("row size: %v\n", row)
// 	fmt.Printf("col size: %v\n", col)
// 	fmt.Printf("vector size: %v\n", vec_size)
// 	assert(vec_size == col, "dimension mismatch")

// 	write_cairo_file(data, dst)
// }

// type Result struct {
// 	Matrix [][]int `json:"matrix"`
// 	Vector []int   `json:"vector"`
// }

// func assert(condition bool, message string) {
// 	if !condition {
// 		fmt.Fprintf(os.Stderr, "Assertion failed: %s\n", message)
// 		os.Exit(1)
// 	}
// }
// func read_json(filename string) Result {
// 	file, err := os.Open(filename) // For read access.
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 	var data Result
// 	decoder := json.NewDecoder(file)
// 	err = decoder.Decode(&data)
// 	if err != nil {
// 		// Handle the error, possibly log it
// 		log.Fatalf("Error decoding JSON: %v", err)
// 	}
// 	// Use the unmarshalled data

// 	return data
// }

// func write_cairo_file(data Result, filename string) {

// 	//opening the file to enter
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		log.Fatalf("Failed to create file: %v", err)
// 	}
// 	defer file.Close()

// 	//opening the data and matrix
// 	var matrix, vector = data.Matrix, data.Vector

// 	// Step 2: Write the header
// 	header := `use array::ArrayTrait;
// use cairomap::matvectmult::{Matrix, Vec, matrixTrait, vecTrait, mapper, reducer, final_output};`
// 	file.WriteString(header)

// 	//Step 3: Write the matrix creation function

// 	// Write the matrix creation function
// 	file.WriteString("//return matrix, row, col\n")
// 	file.WriteString("fn create_matrix() -> (Matrix,u32,u32){\n")
// 	for i, row := range matrix {
// 		file.WriteString(fmt.Sprintf("    let row%d = array![", i+1))
// 		for j, val := range row {
// 			file.WriteString(fmt.Sprintf("%d", val))
// 			if j < len(row)-1 {
// 				file.WriteString(", ")
// 			}
// 		}
// 		file.WriteString("];\n")
// 	}

// 	file.WriteString("    let matrix_array = array![")
// 	for i := range matrix {
// 		file.WriteString(fmt.Sprintf("row%d", i+1))
// 		if i < len(matrix)-1 {
// 			file.WriteString(", ")
// 		}
// 	}
// 	// Dynamically generate the matrix dimensions
// 	file.WriteString("];\n")
// 	rows, cols := len(matrix), len(matrix[0])
// 	file.WriteString(fmt.Sprintf("    let mat = matrixTrait::init_array(%d, %d, @matrix_array);\n", rows, cols))
// 	file.WriteString(fmt.Sprintf("    (mat,%d,%d)\n", rows, cols))
// 	file.WriteString("}\n\n")

// 	// Write the vector creation function
// 	file.WriteString("//return row, vector_length\n")
// 	file.WriteString("fn create_vector()->(Vec,u32){\n")
// 	file.WriteString("    let vec_test = array![")
// 	for i, val := range vector {
// 		file.WriteString(fmt.Sprintf("%d", val))
// 		if i < len(vector)-1 {
// 			file.WriteString(", ")
// 		}
// 	}
// 	// Dynamically generate the vector length
// 	vectorLength := len(vector)
// 	file.WriteString("];\n")
// 	file.WriteString(fmt.Sprintf("    let vec = vecTrait::init_array(%d, @vec_test);\n", vectorLength))
// 	file.WriteString(fmt.Sprintf("    (vec,%d)\n", vectorLength))
// 	file.WriteString("}\n")

// }
