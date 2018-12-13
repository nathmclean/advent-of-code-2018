package main

import (
	"strings"
	"fmt"
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err !=  nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(file)
	ids := []string{}
	for s.Scan() {
		ids = append(ids, s.Text())
	}
	for _, id := range ids {
		for _, id2 := range ids {
			var found int
			var commonLetters []string
			for i, _ := range id {
				if id[i] != id2[i] {
					found++
				} else {
					commonLetters = append(commonLetters, string(id[i]))
				}
			}
			if found == 1 {
				fmt.Println(strings.Join(commonLetters, ""))
				return
			}
		}
	}
}
