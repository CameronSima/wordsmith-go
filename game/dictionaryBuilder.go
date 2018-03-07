package game

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode/utf8"
)

//Build is what builds the dictionary
func (d *Dictionary) Build() {
	allEntries := make(map[string]map[string][]string)
	letters := Points
	for letter := range letters {
		subDict := getSubDict(letter)
		allEntries[letter] = subDict
	}
	d.Entries = allEntries
}

func getSubDict(filename string) map[string][]string {
	subDict := make(map[string][]string)

	absPath, _ := filepath.Abs("assets/" + filename + ".html")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`<P><B>(?P<word>\w+)</B> \(<I>(?P<partOfSpeech>\S+)</I>\) (?P<definition>(?s).*)</P>`)
	for scanner.Scan() {
		match := re.FindStringSubmatch(scanner.Text())

		if len(match) > 2 {
			word := strings.ToUpper(match[1])
			definition := match[3]

			if utf8.RuneCountInString(word) > 2 {
				defs := subDict[word]
				defs = append(defs, definition)
				subDict[word] = defs
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return subDict
}
