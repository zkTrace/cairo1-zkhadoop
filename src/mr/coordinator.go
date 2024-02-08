package mr

// keep this to work with MIT testing framework

import (
	"6.824/mr/coordinator"
)

// MakeCoordinator proxies the call to the new coordinator package.
func MakeCoordinator(files []string, nReduce int) *coordinator.Coordinator {
	return coordinator.MakeCoordinator(files, nReduce)
}
