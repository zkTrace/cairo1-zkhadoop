package coordinator

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"

	"server/mr/common" // Import common types and utilities for MapReduce.
)

// RequestTask is called by workers to request a new map or reduce task.
// It locks the coordinator state, checks the worker's status, and assigns a new task if available.
// If no tasks are available, it sets the reply to indicate if all tasks are done.
func (c *Coordinator) RequestTask(args *RequestTaskArgs, reply *RequestTaskReply) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Check the current status of the worker requesting a task.
	workerStatus := c.getWorkerStatus(args.PID)
	if workerStatus == StatusIdle {
		// Try to assign a map job first.
		if mapJob := c.assignMapJob(args.PID); mapJob != nil {
			reply.MapJob = mapJob
			return nil
		}

		// If no map jobs are available, try to assign a reduce job.
		if reduceJob := c.assignReduceJob(args.PID); reduceJob != nil {
			reply.ReduceJob = reduceJob
			return nil
		}
	}

	// If no tasks are available, check if all jobs are done and set the reply accordingly.
	reply.Done = c.allJobsDone()
	return nil
}

// ReportMapTask is called by workers to report the completion of a map task.
// It updates the task's status to completed, marks the worker as idle, and records the intermediate files produced.
func (c *Coordinator) ReportMapTask(args *ReportMapTaskArgs, reply *ReportMapTaskReply) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	fmt.Println("ReportMapTask")
	// Mark the map task as completed and the worker as idle.
	c.mapStatus[args.InputFile] = JobStatus{StartTime: -1, Status: "completed"}
	c.workerStatus[args.PID] = "idle"

	// Store the intermediate files produced by the map task for later use in reduce tasks.
	for r := 0; r < c.nReducer; r++ {
		c.intermediateFiles[r] = append(c.intermediateFiles[r], args.IntermediateFile[r])
	}
	fmt.Println("End of ReportMapTask")

	return nil
}

// ReportReduceTask is called by workers to report the completion of a reduce task.
// It updates the task's status to completed and marks the worker as idle.
func (c *Coordinator) ReportReduceTask(args *ReportReduceTaskArgs, reply *ReportReduceTaskReply) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Mark the reduce task as completed and the worker as idle.
	c.reduceStatus[args.ReduceNumber] = JobStatus{StartTime: -1, Status: "completed"}
	c.workerStatus[args.PID] = "idle"
	return nil
}

// startServer initializes and starts the RPC server to listen for requests from workers.
// It sets up an HTTP server over a Unix domain socket for inter-process communication.
func (c *Coordinator) startServer() {
	// Register the coordinator's methods as RPC endpoints.
	rpc.Register(c)
	rpc.HandleHTTP()

	// Create a unique socket name for the coordinator and remove any existing socket with the same name.
	sockname := common.CoordinatorSock()
	os.Remove(sockname) // Ignore error as it's fine if the file doesn't exist.

	// Listen on the Unix domain socket.
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	// Serve HTTP over the Unix domain socket in a new goroutine.
	go http.Serve(l, nil)
}
