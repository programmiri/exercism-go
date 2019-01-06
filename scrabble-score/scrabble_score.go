// Package scrabble implements a scrabble game
package scrabble

import "strings"

// Score calculates the scrabble score for a given word
func Score(s string) int {
	letterValue := map[int][]string{
		1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
		2:  {"D", "G"},
		3:  {"B", "C", "M", "P"},
		4:  {"F", "H", "V", "W", "Y"},
		5:  {"K"},
		8:  {"J", "X"},
		10: {"Q", "Z"},
	}
	count := 0
	for k, v := range letterValue {
		currCount := 0
		for i := 0; i < len(v); i++ {
			currScore := strings.Count(strings.ToUpper(s), v[i]) * k
			currCount = currCount + currScore
		}
		count = count + currCount
	}
	return count
}
