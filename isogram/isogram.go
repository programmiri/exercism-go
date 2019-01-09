// Package isogram implements methods to find a isogram.
package isogram

import (
	"regexp"
	"sort"
	"strings"
)

func removeNonAlphanumericChars(s string) string {
	return regexp.MustCompile("[^a-zA-Z0-9]+").ReplaceAllString(s, "")
}

func createdSortedList(s string) []string {
	var list = strings.Split(strings.ToUpper(s), "")
	sort.Strings(list)
	return list
}

// IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(s string) bool {
	onlyChar := removeNonAlphanumericChars(s)
	sortedCharList := createdSortedList(onlyChar)

	if len(sortedCharList) < 1 {
		return true
	}

	result := true
	for i := 0; i < len(sortedCharList)-1; i++ {
		if sortedCharList[i] == sortedCharList[i+1] {
			result = false
			break
		}
	}
	return result
}
