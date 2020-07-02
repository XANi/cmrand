[![GoDoc Reference](https://godoc.org/github.com/XANi/cmrand?status.svg)](http://godoc.org/github.com/XANi/cmrand)
[![Go Report Card](https://goreportcard.com/badge/github.com/XANi/cmrand)](https://goreportcard.com/report/github.com/XANi/cmrand)
https://goreportcard.com/badge/github.com/XANi/cmrand
cmrand
=========

Wrapper providing `math/rand`-compatible interface for calling `crypto/rand`

Usage


```go
package main
import "github.com/XANi/cmrand"
import "fmt"

func main() {
	src := cmrand.New()
	num := src.Intn(12)
	fmt.Println(num)
}
```

 