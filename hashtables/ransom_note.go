package main

import (
	"strings"
)

func RansomNote(magazine, note string) bool {
	magazineMap := freqMap(magazine)
	noteMap := freqMap(note)

	for word, freq := range noteMap {
		if mag, ok := magazineMap[word]; ok {
			if mag < freq {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func freqMap(input string) map[string]int {
	words := strings.Split(input, " ")
	freq := make(map[string]int, len(words)/2)
	for _, word := range words {
		if _, ok := freq[word]; ok {
			freq[word]++
		} else {
			freq[word] = 1
		}
	}
	return freq
}
