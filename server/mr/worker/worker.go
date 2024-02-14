package worker

import (
	"fmt"
)

// Worker continuously requests tasks from the coordinator and processes them.
func Worker(mapf func(string, string) []KeyValue, reducef func(string, []string) string) {
	for {
		taskReply := requestTaskFromCoordinator()
		if taskReply.Done {
			fmt.Println("All jobs done, worker exiting...")
			break
		}

		processTask(taskReply, mapf, reducef)
	}
}
