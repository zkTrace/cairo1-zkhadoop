package main

//
// start the coordinator process, which is implemented
// in ../mr/coordinator.go
//
// go run mrcoordinator.go pg*.txt
//
// Please do not change this file.
//

import (
	"fmt"
	"os"
	"server/mr/common"
	"server/mr/coordinator"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: mrcoordinator inputfiles...\n")
		os.Exit(1)
	}
	
	startTime := time.Now()

	m := coordinator.MakeCoordinator(os.Args[1:], 1)
	for m.Done() == false {
		time.Sleep(time.Second)
	}

	// ====== For HH Demo, verifying list of proofs ======
	fmt.Println("====== Verifying Final 2 Proofs ======") // for now just make the coordiantor the final verifier
	common.VerifyProofs() // for now the proofs in a server/data/mr-tmp/mr-out-0

	// ====== For HH Demo, timing the operation ======
	endTime := time.Now()
  duration := endTime.Sub(startTime)
	fmt.Println("====== This Cairo1 Operation Took: ======") // for now just make the coordiantor the final verifier
	hours := duration / time.Hour
	duration -= hours * time.Hour
	minutes := duration / time.Minute
	duration -= minutes * time.Minute
	seconds := duration / time.Second

	// Print the duration in hours, minutes, and seconds
	fmt.Printf("The operation took %d hours, %d minutes, and %d seconds\n", hours, minutes, seconds)

	time.Sleep(time.Second)
}
