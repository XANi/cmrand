package cmrand

import "fmt"

// Example usage
func Example() {
	src :=New()
	fmt.Println("Hello")
	num := src.Intn(12)
	fmt.Println(num)
	// Output 6
}

