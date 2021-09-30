// Echo program prints all the command line args
package main

import (
	"fmt"
	"os"
)

func main() {
	// Solution -1 can be
	//fmt.Println(os.Args)
	var s, sep string
	for i := 0; i < len(os.Args[1:]); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

}


