package mr

// keep this to work with MIT testing framework

import (
	"6.824/mr/common"
	"6.824/mr/worker"
)

type KeyValue = common.KeyValue

func Worker(mapf func(string, string) []KeyValue,
	reducef func(string, []string) string) {
	worker.Worker(mapf, reducef)
}
