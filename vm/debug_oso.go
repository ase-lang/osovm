package main
import ("fmt"; "os")
func main() {
	fmt.Println("Debug: All arguments:")
	for i, arg := range os.Args {
		fmt.Printf("  Args[%d] = %q\n", i, arg)
	}
	
	if len(os.Args) < 3 {
		fmt.Println("\nUsage: oso run <ritual.oso>")
		os.Exit(1)
	}
	
	// Skip Args[1] if it's a full path duplicate
	startIdx := 1
	if len(os.Args) > 1 && os.Args[1][0] == '/' {
		startIdx = 2
	}
	
	command := os.Args[startIdx]
	ritualFile := os.Args[startIdx+1]
	
	fmt.Printf("\nParsed command: %q\n", command)
	fmt.Printf("Ritual file: %q\n", ritualFile)
}
