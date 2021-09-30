// Compare string concatenation running time
//
// Run with go test -bench=.
package concat_test

import (
	"strings"
	"testing"
)

var args = []string{"hi", "there", "buddy", "boy", "5", "6", "7", "8", "9"}

func concat(args []string) {
	var s, sep string

	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat(args)
	}
}

func BenchmarkJoin(b *testing.B){
	for i := 0; i < b.N; i++ {
		strings.Join(args, " ")
	}
}