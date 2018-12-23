package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type coord struct {
	x, y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var maxX, maxY int
	coords := []coord{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c := coord{}
		fmt.Sscanf(scanner.Text(), "%d, %d", &c.x, &c.y)
		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
		coords = append(coords, c)
	}

	grid := make([][]coord, maxY+1)
	for y := 0; y <= maxY; y++ {
		grid[y] = make([]coord, maxX+1)
		for x := 0; x <= maxX; x++ {
			shortest := -1
			var closest []coord
			for _, c := range coords {
				if x == c.x && y == c.y { // same coord
					closest = []coord{c}
					break
				}
				dist := dist(coord{x, y}, c)
				if shortest < 0 { // first coord
					shortest = dist
					closest = []coord{c}
					continue
				}
				if dist < shortest { // new shortest dist
					shortest = dist
					closest = []coord{c}
					continue
				}
				if dist == shortest { // repeat shortest dist
					shortest = dist
					closest = append(closest, c)
				}
			}
			if len(closest) == 1 {
				grid[y][x] = closest[0]
			} else {
				grid[y][x] = coord{-1, -1}
			}
		}
	}

	// for _, line := range grid {
	// 	fmt.Println(line)
	// }

	infinate := map[coord]bool{}
	for y, col := range grid {
		for x, c := range col {
			if x == 0 || x == maxX {
				infinate[c] = true
				continue
			}
			if y == 0 || y == maxY {
				infinate[c] = true
				continue
			}
		}
	}
	fmt.Println(maxX, maxY)
	fmt.Println(infinate)

	count := map[coord]int{}
	for _, col := range grid {
		for _, c := range col {
			_, ok := infinate[c]
			if ok {
				// fmt.Println(c)
				continue
			}
			count[c]++
		}
	}

	max := 0
	for _, v := range count {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
}

func dist(a, b coord) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}
