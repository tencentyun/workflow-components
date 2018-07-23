package main

import (
	"fmt"
	"os"
)

func main() {
	for i, pair := range os.Environ() {
		fmt.Println(i)
		fmt.Println(pair)
	}
}
