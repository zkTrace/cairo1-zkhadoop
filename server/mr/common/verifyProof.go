package common

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func CollectProofs() []string {
	projectRoot := GetProjectRoot()
	proofDir := filepath.Join(projectRoot, "server/data/mr-tmp")
	proofFiles := []string{} // Initialize an empty slice to store the file names

	err := filepath.Walk(proofDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
					return err // Return error and stop walking the directory
			}
			if !info.IsDir() && filepath.Ext(path) == ".proof" {
					proofFiles = append(proofFiles, info.Name()) // Append the file name to the slice
			}
			return nil
	})

	if err != nil {
			fmt.Printf("Error walking the path %q: %v\n", proofDir, err)
			return nil
	}

	// Process the list of .proof files as needed
	fmt.Println("Proof files found:", proofFiles)
	return proofFiles
}

func VerifyProofs() {
	// Get Proofs
	proofFiles := CollectProofs()
	if len(proofFiles) == 0 {
		log.Fatalf("The proofs array is empty!")
		return 
	}

	// Set up the paths
	projectRoot := GetProjectRoot()
	executionDir := filepath.Join(projectRoot, "server/mr/common") // Updated path
	proofDir := filepath.Join(projectRoot, "server/data/mr-tmp") // Updated path

	// Ensure the output directory exists.
	err := os.MkdirAll(proofDir, 0755) // 0755 is commonly used permission for directories
	if err != nil {
		log.Fatalf("Failed to create directory: %s", err)
	}

	// Loop through
	for _, proofFile := range proofFiles {
		proofFileDir := filepath.Join(proofDir, proofFile)
		// Prepare the command to run the Cairo program. Set the execution and output.
		cmd := exec.Command("bash", "verify_proof.sh", proofFileDir)
		cmd.Dir = executionDir

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		// Run the command.
		err = cmd.Run()
		if err != nil {
			log.Fatalf("Failed to execute command: %s", err)
		}
	}

	fmt.Println("Executed Verify programs successfully")

}