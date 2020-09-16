package main

import (
	"fmt"
	"runtime/debug"

	"github.com/kr/pretty"
)

func main() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		panic("oh dear")
	}
	fmt.Printf("%v\n", pretty.Sprint(bi))
}
