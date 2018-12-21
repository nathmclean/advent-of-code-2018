package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type input struct {
	s       string
	t       time.Time
	message string
}

type ByDate []input

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].t.Before(a[j].t) }

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(file)
	inputs := []input{}
	for s.Scan() {
		input := input{}
		txt := s.Text()
		splitTxt := strings.Split(txt, " ")
		input.s = strings.TrimSuffix(strings.TrimPrefix(fmt.Sprintf("%s %s", splitTxt[0], splitTxt[1]), "["), "]")
		input.message = strings.Join(splitTxt[2:], " ")
		date, err := time.Parse("2006-01-02 15:04", input.s)
		input.t = date
		if err != nil {
			log.Fatal(err)
		}
		inputs = append(inputs, input)
	}
	dateSortedInputs := make(ByDate, len(inputs))
	for i, input := range inputs {
		dateSortedInputs[i] = input
	}
	sort.Sort(dateSortedInputs)

	type guard struct {
		sleepTime   int
		sleepMinute []int
	}

	guards := map[string]guard{}
	currentGuard := ""
	asleepmin := 0
	for _, d := range dateSortedInputs {
		startRegex := regexp.MustCompile("begins shift")
		if startRegex.Match([]byte(d.message)) {
			_, err = fmt.Sscanf(d.message, "Guard #%s begins shift", &currentGuard)
			if err != nil {
				log.Fatal(err)
			}
			_, ok := guards[currentGuard]
			if !ok {
				guards[currentGuard] = guard{0, make([]int, 60)}
			}
		}
		if d.message == "falls asleep" {
			asleepmin = d.t.Minute()
		}
		guard := guards[currentGuard]
		if d.message == "wakes up" {
			for i := asleepmin; i < d.t.Minute(); i++ {
				guard.sleepMinute[i]++
				guard.sleepTime++
			}
		}
		guards[currentGuard] = guard
	}

	mostSleepyGuardtime := 0
	mostSleepyGuard := ""
	for guardId, guard := range guards {
		if guard.sleepTime > mostSleepyGuardtime {
			mostSleepyGuardtime = guard.sleepTime
			mostSleepyGuard = guardId
		}
	}

	fmt.Println(mostSleepyGuard)

	minute := 0
	max := guards[mostSleepyGuard].sleepMinute[0]
	for i, v := range guards[mostSleepyGuard].sleepMinute {
		if v > max {
			minute = i
			max = v
		}
	}
	fmt.Println(minute)
	idNumber, err := strconv.Atoi(mostSleepyGuard)
	if err !=  nil {
		log.Fatal(err)
	}
	fmt.Println(idNumber * minute)
}
