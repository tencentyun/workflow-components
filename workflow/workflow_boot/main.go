package main

import (
	"fmt"
	"os"
)

func main() {

	builder, err := NewBuilder()
	if err != nil {
		fmt.Println("BUILDER FAILED: ", err)
		os.Exit(1)
	}

	if err := builder.run(); err != nil {
		fmt.Println("INIT FAILED: ", err)
		os.Exit(1)
	} else {
		fmt.Println("INIT SUCCEED")
	}
}
