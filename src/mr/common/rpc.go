package common

import (
	"fmt"
	"os"
)

// MapJob represents a mapping task with input file, job number, and the count of reducers.
type MapJob struct {
	InputFile    string
	MapJobNumber int
	ReducerCount int
}

// ReduceJob represents a reducing task with intermediate files and a reducer number.
type ReduceJob struct {
	IntermediateFiles []string
	ReduceNumber      int
}

// RequestTaskArgs holds arguments for requesting a task, including the process ID.
type RequestTaskArgs struct {
	PID int // PID is the process ID.
}

// RequestTaskReply defines the reply structure for a task request, including the task details and completion status.
type RequestTaskReply struct {
	MapJob    *MapJob
	ReduceJob *ReduceJob
	Done      bool
}

// ReportMapTaskArgs holds arguments for reporting a completed map task, including the input file, intermediate files, and process ID.
type ReportMapTaskArgs struct {
	InputFile        string
	IntermediateFile []string
	PID              int // PID is the process ID.
}

// ReportMapTaskReply defines the reply structure for reporting a completed map task.
type ReportMapTaskReply struct {
}

// ReportReduceTaskArgs holds arguments for reporting a completed reduce task, including the process ID and reducer number.
type ReportReduceTaskArgs struct {
	PID          int // PID is the process ID.
	ReduceNumber int
}

// ReportReduceTaskReply defines the reply structure for reporting a completed reduce task.
type ReportReduceTaskReply struct {
}

// CoordinatorSock generates a unique-ish UNIX-domain socket name in /var/tmp for the coordinator, avoiding issues with directories that do not support UNIX-domain sockets.
func CoordinatorSock() string {
	return fmt.Sprintf("/var/tmp/824-mr-%d", os.Getuid())
}
