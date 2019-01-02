//Package raindrops adds a method to translate numbers in raindrop-speak
package raindrops

import (
	"strconv"
	"strings"
)

// Convert converts a number to a string in raindrop-speak
func Convert(num int) string {
	var r []string
	if num%3 == 0 {
		r = append(r, "Pling")
	}
	if num%5 == 0 {
		r = append(r, "Plang")
	}
	if num%7 == 0 {
		r = append(r, "Plong")
	}

	if len(r) == 0 {
		r = append(r, strconv.Itoa(num))
	}
	return strings.Join(r, "")
}
