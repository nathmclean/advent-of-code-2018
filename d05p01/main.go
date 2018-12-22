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
	length := 0
	for {
		for i := 0; i < len(input); i++ {
			if i == len(input)-1 {
				continue
			}
			if pair(input[i], input[i+1]) {
				input = append(input[:i], input[i+2:]...)
				break
			}
		}
		if length == len(input) {
			break
		}
		length = len(input)
	}
	fmt.Println(length)
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
