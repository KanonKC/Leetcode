package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
    wordList := strings.Split(s, " ")
	lastWord := ""

	for _, word := range wordList {
		if word == "" {
			continue
		}
		lastWord = word
	}

	return len(lastWord)
}

func main() {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
}