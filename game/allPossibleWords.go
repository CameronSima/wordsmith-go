package game

import (
	"strings"
)

// GetAllPossibleWords from letterset
func GetAllPossibleWords(letters []Letter, dictionary *Dictionary) map[string][]string {
	words := make(map[string][]string)
	letterMap := letterToCountMap(letters)

	for _, subDict := range dictionary.Entries {
		for word, definitions := range subDict {
			if canBuildWord(letterMap, word) == true {
				words[word] = definitions
			}
		}
	}
	return words
}

func canBuildWord(letters map[string]int, word string) bool {
	result := true
	wordLettArr := strings.Split(word, "")
	wordMap := toCountMap(wordLettArr)

	for letter, count := range wordMap {
		count2 := letters[letter]
		if count > count2 {
			result = false
		}
	}
	return result
}

func toCountMap(arr []string) map[string]int {
	result := make(map[string]int)
	for _, lett := range arr {
		if _, exists := result[lett]; exists {
			result[lett]++
		} else {
			result[lett] = 1
		}
	}
	return result
}

func letterToCountMap(arr []Letter) map[string]int {
	result := make(map[string]int)
	for _, lett := range arr {
		if _, exists := result[lett.Value]; exists {
			result[lett.Value]++
		} else {
			result[lett.Value] = 1
		}
	}
	return result
}
