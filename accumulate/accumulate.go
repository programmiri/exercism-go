// Package accumulate implements a method to accumulate collections
package accumulate

// import (
// 	"fmt"
// 	"strings"
// 	"testing"
// )

// Accumulate takes a collection, performs an operation on this
// collection and returns the result
func Accumulate(collection []string, operation func(string) string) []string {
	var acc = []string{}
	for _, el := range collection {
		acc = append(acc, operation(el))
	}
	return acc
}
