package common

type KeyValue struct {
	Key   string
	Value string
}

// Helper function to get the project root from the environment variable
func GetProjectRoot() string {
	// os.Clearenv()
	// os.Setenv("PROJECT_ROOT","/home/ec2-user/zkhadoop-cairo1")  

	// projectRoot := os.Getenv("PROJECT_ROOT")
	// log.Printf("GOT PROJECT ROOT: " + projectRoot)
	// if projectRoot == "" {
	// 	log.Fatal("PROJECT_ROOT environment variable is not set.")
	// }
	return "/home/ec2-user/zkhadoop-cairo1/"
}
