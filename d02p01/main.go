package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err !=  nil {
		log.Fatal(err)
	}

	var twos, threes int
	s := bufio.NewScanner(file)
	for s.Scan() {
		counts := map[rune]int{}
		id := s.Text()
		for _, r := range id {
			counts[r]++
		}
		var two, three bool
		for _, v := range counts {
			if v == 2 && !two {
				twos ++
				two = true
			}
			if v == 3 && !three  {
				threes++
				three = true
			}
		}
	}
	log.Println(twos * threes)
}
