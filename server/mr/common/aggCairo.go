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