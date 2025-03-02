package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {

	charMap := make(map[byte]byte)
	inverseCharMap := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if (charMap[s[i]] == 0) {
			if inverseCharMap[t[i]] == 0 {
				charMap[s[i]] = t[i]
				inverseCharMap[t[i]] = s[i]
			} else if inverseCharMap[t[i]] != s[i] {
				return false
			}
		} else if charMap[s[i]] != t[i] {
			return false
		}
	}

    return true
}

func main() {
	fmt.Println(isIsomorphic("egg","add"))
	fmt.Println(isIsomorphic("foo","bar"))
	fmt.Println(isIsomorphic("paper","title"))
	fmt.Println(isIsomorphic("badc","pzpz"))
}