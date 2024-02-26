package worker

import (
	"hash/fnv"
	"server/mr/common"
)

type KeyValue = common.KeyValue
type MapJob = common.MapJob
type ReduceJob = common.ReduceJob
type RequestTaskArgs = common.RequestTaskArgs
type RequestTaskReply = common.RequestTaskReply
type ReportMapTaskArgs = common.ReportMapTaskArgs
type ReportReduceTaskReply = common.ReportReduceTaskReply
type ReportReduceTaskArgs = common.ReportReduceTaskArgs
type ReportMapTaskReply = common.ReportMapTaskReply

// ByKey implements sort.Interface for []KeyValue based on the Key field.
type ByKey []KeyValue

func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

// ihash hashes a string to a non-negative integer.
func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

// partitionByKey distributes key-value pairs among reducers based on their keys.
func partitionByKey(kva []KeyValue, reduceCount int) [][]KeyValue {
	partitionedKva := make([][]KeyValue, reduceCount)
	for _, kv := range kva {
		partition := ihash(kv.Key) % reduceCount
		partitionedKva[partition] = append(partitionedKva[partition], kv)
	}
	return partitionedKva
}
