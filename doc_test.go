package cmrand

import "fmt"

// Example usage
func Example() {
	src := cmrand.New()
	num := src.Intn(12)
	fmt.Println(num)
	// Output: 6
}

