package coordinator

import (
	"sync"

	"6.824/mr/common" // Importing common types and constants used in the MapReduce framework
)

// Type aliases for request and reply structs for tasks and job reports.
// This improves readability and allows for easy updates if the underlying types change.
type RequestTaskArgs = common.RequestTaskArgs
type RequestTaskReply = common.RequestTaskReply
type ReportMapTaskArgs = common.ReportMapTaskArgs
type ReportMapTaskReply = common.ReportMapTaskReply
type ReportReduceTaskArgs = common.ReportReduceTaskArgs
type ReportReduceTaskReply = common.ReportReduceTaskReply
type MapJob = common.MapJob
type ReduceJob = common.ReduceJob

// Constants defining possible statuses of tasks and workers.
const (
	StatusIdle      = "idle"
	StatusBusy      = "busy"
	StatusCompleted = "completed"
	StatusPending   = "pending"
	StatusRunning   = "running"
)

// Coordinator struct represents the central coordinator of the MapReduce job.
// It keeps track of worker statuses, job statuses, and manages task assignments.
type Coordinator struct {
	workerStatus      map[int]string       // Worker status indexed by worker ID
	mapStatus         map[string]JobStatus // Status of map tasks by file name
	mapTaskNumber     int                  // Counter to keep track of assigned map tasks
	reduceStatus      map[int]JobStatus    // Status of reduce tasks by reduce task number
	nReducer          int                  // Total number of reducers
	intermediateFiles map[int][]string     // Mapping from reduce task number to intermediate file names
	mu                sync.Mutex           // Mutex for concurrent access to the Coordinator's fields
}

// JobStatus struct stores the start time and current status of a job (map or reduce task).
type JobStatus struct {
	StartTime int64  // Start time of the task in Unix timestamp
	Status    string // Current status of the task
}

// MakeCoordinator initializes a new Coordinator with a given set of input files and a number of reducers.
// It prepares the map and reduce tasks and starts the server and ticker for task management.
func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{
		mapTaskNumber:     0,
		nReducer:          nReduce,
		workerStatus:      make(map[int]string),
		mapStatus:         make(map[string]JobStatus),
		reduceStatus:      make(map[int]JobStatus),
		intermediateFiles: make(map[int][]string),
	}

	// Initialize map tasks as pending for each input file
	for _, file := range files {
		c.mapStatus[file] = JobStatus{Status: StatusPending}
	}

	// Initialize reduce tasks as pending
	for i := 0; i < nReduce; i++ {
		c.reduceStatus[i] = JobStatus{Status: StatusPending}
	}

	c.startTicker() // Start a ticker to periodically check for dead workers
	c.startServer() // Start the RPC server to listen for worker requests
	return &c
}
