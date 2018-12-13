package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		var n int
		_, err = fmt.Sscanf(scanner.Text(), "%d", &n)
		if err != nil {
			log.Fatal(err)
		}
		sum += n
	}
	log.Println(sum)
}
