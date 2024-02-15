package worker

import (
	"fmt"
	"log"
	"net/rpc"
	"os"

	"server/mr/common"
)

// requestTaskFromCoordinator sends an RPC request to the coordinator to fetch a new map or reduce task.
// It constructs the request with the current process ID to uniquely identify the worker.
// The function returns the task details provided by the coordinator.
func requestTaskFromCoordinator() RequestTaskReply {
	args := RequestTaskArgs{PID: os.Getpid()}      // Prepare the request with the worker's PID.
	var reply RequestTaskReply                     // Holder for the coordinator's response.
	call("Coordinator.RequestTask", &args, &reply) // Send the RPC request.
	return reply                                   // Return the received task information.
}

// reportMapTaskToCoordinator sends an RPC request to the coordinator to report the completion of a map task.
// It includes details about the input file and the generated intermediate files, along with the worker's PID.
func reportMapTaskToCoordinator(inputFile string, intermediateFiles []string) {
	args := ReportMapTaskArgs{
		InputFile:        inputFile,         // The input file processed by the map task.
		IntermediateFile: intermediateFiles, // The output intermediate files produced by the map task.
		PID:              os.Getpid(),       // The worker's PID.
	}
	var reply ReportMapTaskReply // Holder for any response, though typically not used here.
	fmt.Println("in report reportMapTaskToCoordinator", args)
	call("Coordinator.ReportMapTask", &args, &reply) // Send the completion report.
}

// reportReduceTaskToCoordinator sends an RPC request to the coordinator to report the completion of a reduce task.
// It includes the reduce task number and the worker's PID.
func reportReduceTaskToCoordinator(reduceNumber int) {
	args := ReportReduceTaskArgs{
		ReduceNumber: reduceNumber, // The number of the reduce task that was completed.
		PID:          os.Getpid(),  // The worker's PID.
	}
	var reply ReportReduceTaskReply                     // Holder for any response, typically not used here.
	call("Coordinator.ReportReduceTask", &args, &reply) // Send the completion report.
}

// call performs the actual RPC call to the coordinator.
// It establishes a connection to the coordinator via a Unix socket, sends the request, and waits for the reply.
// If the RPC call fails, it logs a fatal error, terminating the worker process.
func call(rpcName string, args interface{}, reply interface{}) {
	sockName := common.CoordinatorSock()          // Get the Unix socket name for connecting to the coordinator.
	client, err := rpc.DialHTTP("unix", sockName) // Establish the RPC connection.
	if err != nil {
		log.Fatal("dialing:", err) // Fatal log if connection fails.
	}
	defer client.Close() // Ensure the connection is closed after the call.

	// Make the RPC call with the specified method name, arguments, and a place to store the reply.
	if err := client.Call(rpcName, args, reply); err != nil {
		log.Fatalf("%s call failed: %v", rpcName, err) // Log fatal error if the RPC call fails.
	}
}
