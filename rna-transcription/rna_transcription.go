// Package strand implements a method to extrakt information from a DNA strand
package strand

import "strings"

// ToRNA return its RNA complement from a given DNA strand
func ToRNA(dna string) string {
	r := strings.NewReplacer("G", "C",
		"C", "G",
		"T", "A",
		"A", "U")
	return r.Replace(dna)

}
