// Package dna implements a method to extrakt information from a DNA strand
package dna

import (
	"errors"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

func validNuc(d DNA) bool {
	for _, r := range d {
		if r != 'A' && r != 'C' && r != 'G' && r != 'T' {
			return false
		}
	}
	return true

}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	var h Histogram
	if len(d) > 0 && !validNuc(d) {
		return h, errors.New("ivalid nucleotides")
	}
	nucs := [4]rune{'A', 'C', 'G', 'T'}
	h = make(map[rune]int)
	for _, r := range nucs {

		h[r] = strings.Count(string(d), string(r))
	}
	return h, nil
}

//
