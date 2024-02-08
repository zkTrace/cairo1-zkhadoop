package worker

import (
	"log"
	"os"
	"sort"
)

// processTask dispatches the task to either map or reduce processing.
// TODO: Cairo code generation -> message coordinator with stack trace
func processTask(reply RequestTaskReply, mapf func(string, string) []KeyValue, reducef func(string, []string) string) {
	if reply.MapJob != nil {
		processMapTask(reply.MapJob, mapf)
	} else if reply.ReduceJob != nil {
		processReduceTask(reply.ReduceJob, reducef)
	}
}

// processMapTask handles the map task, including reading input, executing the map function, and storing the output.
func processMapTask(job *MapJob, mapf func(string, string) []KeyValue) {
	content, err := os.ReadFile(job.InputFile)
	if err != nil {
		log.Fatalf("cannot read %v", job.InputFile)
	}

	kva := mapf(job.InputFile, string(content))
	sort.Sort(ByKey(kva))

	partitionedKva := partitionByKey(kva, job.ReducerCount)
	intermediateFiles := writeIntermediateFiles(partitionedKva, job.MapJobNumber)

	reportMapTaskToCoordinator(job.InputFile, intermediateFiles)
}

// processReduceTask handles the reduce task, including reading intermediate files, executing the reduce function, and writing the output.
func processReduceTask(job *ReduceJob, reducef func(string, []string) string) {
	intermediate := readIntermediateFiles(job.IntermediateFiles)
	sort.Sort(ByKey(intermediate))

	writeReduceOutput(intermediate, job.ReduceNumber, reducef)
	reportReduceTaskToCoordinator(job.ReduceNumber)
}
