package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	raw, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	input := []rune(string(raw))
	seen := map[rune]bool{}
	for _, v := range input {
		char := unicode.ToLower(v)
		seen[char] = true
	}

	minLength := len(input)
	for char := range seen {
		newInput := []rune{}
		for _, v := range input {
			if unicode.ToLower(v) != char {
				newInput = append(newInput, v)
			}
		}
		length := 0
		for {
			for i := 0; i < len(newInput); i++ {
				if i == len(newInput)-1 {
					continue
				}
				if pair(newInput[i], newInput[i+1]) {
					newInput = append(newInput[:i], newInput[i+2:]...)
					break
				}
			}
			if length == len(newInput) {
				break
			}
			length = len(newInput)
		}
		if length < minLength {
			minLength = length
		}
	}
	fmt.Println(minLength)
}

func pair(a, b rune) bool {
	if unicode.IsUpper(a) {
		if unicode.ToLower(a) != b {
			return false
		}
	} else {
		if unicode.ToUpper(a) != b {
			return false
		}
	}
	return true
}
