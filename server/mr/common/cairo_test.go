package common

import (
	"testing"
)

func TestInputToCairo(t *testing.T) {
	// Test if I can convert a input json into cairo code
	t.Log("Starting file generatorion test for input file")
	var filename = "../../data/testing/input.json"
	var dst = "../../data/testing/matvecdata_mapper.cairo"
	ConvertJsonToCairo(filename, dst)
}

func TestMapping(t *testing.T) {
	// Test if I can get a mapper result
	t.Log("Starting Mapping Test")
	var dst = "../../data/testing/mr"
	CalLCairoMap(1, dst)
}

func TestIntermediaryToCairo(t *testing.T) {
	// your test code here
	t.Log("Starting file generation test for intermediary file")
	var input = "../../data/testing/mr/mr-1-1"
	var dst = "../../data/testing/matvecdata_reducer.cairo"
	ConvertIntermediateToCairo(input, dst)
}

func TestCallingReduce(t *testing.T) {
	// Testing if the resulting file is correct
	t.Log("Output final reduce file")
	var dst = "../../data/testing/reduce"
	var jobid string = "1"
	CallCairoReduce(jobid, dst)
}
