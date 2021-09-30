// Echo program prints all the command line args
package main

import (
	"fmt"
	"os"
)

func main() {
	/*
	for i := 0; i < len(os.Args[1:]); i++ {
		fmt.Println(i, os.Args[i])
	}*/
	for i, arg := range os.Args {
		fmt.Printf("%d: %s\n", i, arg)
	}
}

