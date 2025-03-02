package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	charMap := make(map[rune]int)
	for _, t := range magazine {
		charMap[t]++
	}

	for _, t := range ransomNote {
		if charMap[t] == 0 {
			return false
		}
		charMap[t]--
	}
    return true
}

func main() {
	fmt.Println(canConstruct("a","b"))
	fmt.Println(canConstruct("aa","aab"))
}