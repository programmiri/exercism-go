// Package proverb implements a method to create a proverb
package proverb

// Proverb takes a input and creates a proverb in a special manner with it.
func Proverb(rhyme []string) []string {
	var s = []string{}

	if len(rhyme) == 0 {
		return s
	}

	for i := 0; i < len(rhyme)-1; i++ {
		s = append(s, "For want of a "+rhyme[i]+" the "+rhyme[i+1]+" was lost.")
	}
	return append(s, "And all for the want of a "+rhyme[0]+".")

}
