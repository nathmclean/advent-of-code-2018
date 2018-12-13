package main

import (
	"strconv"
	"strings"
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
	cloth := make([][]int, 1000)
	for i := range cloth {
		cloth[i] = make([]int, 1000)
	}
	for s.Scan() {
		splitText := strings.Split(s.Text(), " ")
		positions :=  strings.Split(splitText[2], ",")
		dimentions := strings.Split(splitText[3], "x")
		top, err := strconv.Atoi(positions[0])
		if err !=  nil {
			log.Fatal(err)
		}
		left, err := strconv.Atoi(strings.TrimSuffix(positions[1], ":"))
		if err !=  nil {
			log.Fatal(err)
		}
		tall, err := strconv.Atoi(dimentions[0])
		if err !=  nil {
			log.Fatal(err)
		}
		wide, err  := strconv.Atoi(dimentions[1])
		if err !=  nil {
			log.Fatal(err)
		}
		
		for y := top; y < top + tall; y++ {
			for x := left; x < left + wide ; x++ {
				cloth[y][x]++
			}
		}
	}

	dup := 0
	for y := range cloth {
		for x := range cloth {
			if cloth[y][x] > 1 {
				dup ++
			}
		}
	}
	log.Println(dup)
}
