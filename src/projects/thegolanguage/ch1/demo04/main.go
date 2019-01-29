package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[0:] {
		fmt.Print(index)
		fmt.Println(" value=" + arg)
	}
}
