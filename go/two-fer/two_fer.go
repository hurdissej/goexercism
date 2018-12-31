// Package twofer implements a simple library for calculating two for ones
package twofer

import "fmt"

// ShareWith returns a string of who this is being shared with
func ShareWith(name string) string {
	var result string
	if len(name) > 0 {
		result = name
	} else {
		result = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", result)
}
