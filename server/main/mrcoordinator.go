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
	"time"

	"server/mr/common"
	"server/mr/coordinator"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: mrcoordinator inputfiles...\n")
		os.Exit(1)
	}

	m := coordinator.MakeCoordinator(os.Args[1:], 1)
	for m.Done() == false {
		time.Sleep(time.Second)
	}

	// ====== For HH Demo, verifying list of proofs ======
	fmt.Println("====== Verifying Final 2 Proofs ======") // for now just make the coordiantor the final verifier
	common.VerifyProofs() // for now the proofs in a server/data/mr-tmp/mr-out-0

	time.Sleep(time.Second)
}
