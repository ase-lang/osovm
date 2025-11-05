package main
import ("fmt"; "os")
func main() {
	fmt.Printf("Num args: %d\n", len(os.Args))
	for i, arg := range os.Args {
		fmt.Printf("Args[%d] = %q\n", i, arg)
	}
}
