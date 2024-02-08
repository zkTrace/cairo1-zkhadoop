package coordinator

import (
	"time"
)

// assignMapJob attempts to assign a pending map job to a worker.
// It iterates over all map tasks and assigns the first pending task to the worker.
// Once assigned, it updates the task's status to running and the worker's status to busy.
func (c *Coordinator) assignMapJob(workerID int) *MapJob {
	for filename, status := range c.mapStatus {
		if status.Status == StatusPending {
			job := &MapJob{
				InputFile:    filename,
				MapJobNumber: c.mapTaskNumber,
				ReducerCount: c.nReducer,
			}
			// Update the status of the map task and worker
			c.mapStatus[filename] = JobStatus{StartTime: time.Now().Unix(), Status: StatusRunning}
			c.workerStatus[workerID] = StatusBusy
			c.mapTaskNumber++
			return job
		}
	}
	return nil // No pending map jobs available
}

// assignReduceJob attempts to assign a pending reduce job to a worker.
// It first checks if all map jobs are completed. If so, it assigns the first pending reduce task.
// The task's status is updated to running, and the worker's status to busy.
func (c *Coordinator) assignReduceJob(workerID int) *ReduceJob {
	if !c.allMapJobDone() { // Ensure all map jobs are completed before assigning reduce jobs
		return nil
	}

	for i, status := range c.reduceStatus {
		if status.Status == StatusPending {
			job := &ReduceJob{
				IntermediateFiles: c.intermediateFiles[i],
				ReduceNumber:      i,
			}
			// Update the status of the reduce task and worker
			c.reduceStatus[i] = JobStatus{StartTime: time.Now().Unix(), Status: StatusRunning}
			c.workerStatus[workerID] = StatusBusy
			return job
		}
	}
	return nil // No pending reduce jobs available
}

// allMapJobDone checks if all map jobs have been completed.
func (c *Coordinator) allMapJobDone() bool {
	for _, v := range c.mapStatus {
		if v.Status != "completed" {
			return false // At least one map job is not yet completed
		}
	}
	return true // All map jobs are completed
}

// getWorkerStatus retrieves the current status of a worker by its ID.
// If no status is found, it defaults to StatusIdle and updates the worker's status.
func (c *Coordinator) getWorkerStatus(workerID int) string {
	if status, exists := c.workerStatus[workerID]; exists {
		return status
	}
	c.workerStatus[workerID] = StatusIdle
	return StatusIdle
}

// allJobsDone checks if all map and reduce jobs have been completed.
func (c *Coordinator) allJobsDone() bool {
	for _, status := range c.mapStatus {
		if status.Status != StatusCompleted {
			return false // A map job is not yet completed
		}
	}
	for _, status := range c.reduceStatus {
		if status.Status != StatusCompleted {
			return false // A reduce job is not yet completed
		}
	}
	return true // All jobs are completed
}

// checkDeadWorkers checks for workers that have not updated their status in a timely manner.
// It resets any tasks assigned to dead workers back to pending status.
func (c *Coordinator) checkDeadWorkers() {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now().Unix()
	for k, v := range c.mapStatus {
		if v.Status == "running" && now > (v.StartTime+10) {
			c.mapStatus[k] = JobStatus{StartTime: -1, Status: "pending"}
		}
	}

	for k, v := range c.reduceStatus {
		if v.Status == "running" && now > (v.StartTime+10) {
			c.reduceStatus[k] = JobStatus{StartTime: -1, Status: "pending"}
		}
	}
}

// startTicker initializes a ticker that triggers every 10 seconds to check for dead workers.
func (c *Coordinator) startTicker() {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			if c.Done() {
				ticker.Stop()
				return // Stop the ticker if all jobs are done
			}
			c.checkDeadWorkers()
		}
	}()
}

// Done checks if the entire MapReduce job has finished, specifically if all reduce jobs are completed.
func (c *Coordinator) Done() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, v := range c.reduceStatus {
		if v.Status != "completed" {
			return false // At least one reduce job is not yet completed
		}
	}
	return true // All reduce jobs are completed
}
