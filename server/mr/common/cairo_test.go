package common

import (
	"testing"
)

func TestCollectProofs(t *testing.T) {
	// your test code here
	t.Log("Starting proof colleciton from /app/server/data/mr-tmp")
	CollectProofs()
}
func TestAggregateCairo(t *testing.T) {
	// your test code here
	t.Log("Starting file generation test for aggregated mapper file")
	// aggMapDst := "/app/cairo/map/src/agg-lib.cairo"
	aggRedDst := "/app/cairo/reducer/src/agg-lib.cairo"
	// AggregateMapperCairo(aggMapDst)
	AggregateReducerCairo(aggRedDst)
}

func TestInputToCairo(t *testing.T) {
	// Test if I can convert a input json into cairo code
	t.Log("Starting file generatorion test for input file")
	var filename = "/app/server/data/testing/input.json"
	var dst = "/app/cairo/map/src/matvecdata_mapper.cairo"
	ConvertJsonToCairo(filename, dst)
}

func TestMapping(t *testing.T) {
	// Test if I can get a mapper result
	t.Log("Starting Mapping Test")
	var dst = "/app/server/data/testing/mr"
	CallCairoMap(1, dst)
}

func TestIntermediaryToCairo(t *testing.T) {
	// your test code here
	t.Log("Starting file generation test for intermediary file")
	var input = "/app/server/data/testing/mr/mr-1-1"
	var dst = "/app/cairo/reducer/src/matvecdata_reducer.cairo"
	ConvertIntermediateToCairo(input, dst)
}

func TestCallingReduce(t *testing.T) {
	// Testing if the resulting file is correct
	t.Log("Output final reduce file")
	var dst = "/app/server/data/testing/reduce"
	var jobid string = "1"
	CallCairoReduce(jobid, dst)
}
