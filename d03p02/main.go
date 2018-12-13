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

	s := bufio.NewScanner(file)
	var cloth [1100][1100]int
	type claim struct {
		id int
		y int
		x int
		h int
		w int
	}
	claims := []*claim{}
	for s.Scan() {
		c := &claim{}
		fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &c.id, &c.x, &c.y, &c.w, &c.h)
		claims = append(claims, c)
	}

	for _, c := range claims {
		for x := c.x; x < c.x + c.w; x++ {
			for y := c.y; y < c.y + c.h; y++ {
				cloth[x][y]++
			}
		}
	}

	log.Println("look for valid Claim")
	CLAIM:
	for _, c := range claims {
			for x := c.x; x < c.x + c.w; x++ {
				for y := c.y; y < c.y + c.h; y++ {
					fmt.Println("continue")
					if cloth[x][y] != 1 {continue CLAIM}
				}
			}
		log.Println(c.id)
	}
	log.Println("done")
}