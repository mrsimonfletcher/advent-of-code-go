package main

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func excludesBadSubstrings(str string) bool {
	return !regexp.MustCompile("ab|cd|pq|xy").MatchString(str)
}

func threeOrMoreVowels(str string) bool {
	vowels := []string{"a", "e", "i", "o", "u"}
	count := 0

	for _, vowel := range vowels {
		count += strings.Count(str, vowel)
	}

	return count >= 3
}

func repeatingLetter(str string) bool {
	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			return true
		}
	}

	return false
}

func goodString(str string) bool {
	return excludesBadSubstrings(str) &&
		threeOrMoreVowels(str) &&
		repeatingLetter(str)
}

func main() {
	input, err := ioutil.ReadFile("./day-05-input.txt")
	if err != nil {
		panic(err)
	}

	strings := strings.Split(string(input), "\n")
	count := 0

	for _, str := range strings {
		if goodString(str) {
			count++
		}
	}

	println(count)
}
