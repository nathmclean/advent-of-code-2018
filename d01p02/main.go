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
	nums := []int{}
	for scanner.Scan() {
		var n int
		_, err = fmt.Sscanf(scanner.Text(), "%d", &n)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, n)
	}

	seen := map[int]bool{0: true}
	sum := 0
	for {
		for _, n := range nums {
			sum += n
			log.Println("sum", n, sum)
			_, ok := seen[sum]
			if ok {
				log.Println(sum)
				return
			}
			seen[sum] = true
		}
	}
}
