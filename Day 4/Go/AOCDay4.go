package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("day4.dat")
	check(err)

	lines := strings.Split(string(dat), "\r\n")
	sortedEntries := part1(lines)

	/*f, err := os.Create("orderedDat.dat")
	check(err)
	defer f.Close()
	for i := 0; i < len(sortedEntries); i++ {
		f.WriteString(sortedEntries[i].entryTime.Format("[2006-01-02 15:04]") + " " + sortedEntries[i].action + "\r\n")
	}
	f.Sync()*/

	guardSleepLog(sortedEntries)
}

type entry struct {
	entryTime time.Time
	action    string
}

type sortEntryByTime []entry

func (e sortEntryByTime) Len() int           { return len(e) }
func (e sortEntryByTime) Less(i, j int) bool { return e[i].entryTime.Before(e[j].entryTime) }
func (e sortEntryByTime) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

type sleepTime struct {
	start int
	stop  int
}

type guardSleep struct {
	id       int
	sleepLog []sleepTime
}

func (s sleepTime) Total() int { return s.stop - s.start }

func (g guardSleep) Total() int {
	t := 0
	for i := 0; i < len(g.sleepLog); i++ {
		t += g.sleepLog[i].Total()
	}
	return t
}

func part1(lines []string) []entry {
	entries := []entry{}

	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], "]")

		t, _ := time.Parse("[2006-01-02 15:04", parts[0])
		e := entry{entryTime: t, action: strings.TrimPrefix(parts[1], " ")}
		entries = append(entries, e)
	}
	sort.Sort(sortEntryByTime(entries))

	return entries
}

func guardSleepLog(entries []entry) {
	guardMap := make(map[int]guardSleep)

	var guardEntry guardSleep
	sleepStart := -1
	for i := 0; i < len(entries); i++ {
		if strings.HasPrefix(entries[i].action, "Guard #") {
			if guardEntry.id != 0 {
				addToGuards(guardMap, guardEntry)
			}
			parts := strings.Split(entries[i].action, " ")
			id, _ := strconv.Atoi(strings.TrimPrefix(parts[1], "#"))

			guardEntry = guardSleep{id: id}
			//guardEntry.sleepLog := []sleepTime)
		} else if strings.HasPrefix(entries[i].action, "falls asleep") {
			sleepStart = entries[i].entryTime.Minute()

		} else if strings.HasPrefix(entries[i].action, "wakes up") {
			sleep := sleepTime{start: sleepStart, stop: entries[i].entryTime.Minute()}
			guardEntry.sleepLog = append(guardEntry.sleepLog, sleep)
		}
	}
	addToGuards(guardMap, guardEntry)

	fmt.Println(guardMap)

	maxID := 0
	max := 0
	for _, v := range guardMap {
		fmt.Println("ID ", v.id, " = ", v.Total())
		if v.Total() > max {
			max = v.Total()
			maxID = v.id
		}
	}
	fmt.Println("** max id = ", maxID, " with ", max, " minutes")

	guard := guardMap[maxID]
	minute, _ := findMostCommonMinute(guard)
	fmt.Println("Most common minute for max ", minute)
	fmt.Println("Solution = ", (maxID * minute))
	id, min := findMostCommonMinuteInMap(guardMap)
	fmt.Println("Most common minute for all ", id, " = ", min)
	fmt.Println("Solution part 2 = ", (id * min))

}

func addToGuards(guardMap map[int]guardSleep, guard guardSleep) {
	if val, ok := guardMap[guard.id]; ok {
		for i := 0; i < len(guard.sleepLog); i++ {
			val.sleepLog = append(val.sleepLog, guard.sleepLog[i])
		}
		guardMap[guard.id] = val
	} else {
		guardMap[guard.id] = guard
	}
}

func findMostCommonMinuteInMap(guardMap map[int]guardSleep) (int, int) {
	largestID := -1
	largestMinute := -1
	max := -1
	for k, v := range guardMap {
		m, c := findMostCommonMinute(v)

		if c > max {
			max = c
			largestMinute = m
			largestID = k
		}
	}
	return largestID, largestMinute
}

func findMostCommonMinute(guard guardSleep) (int, int) {
	minuteMap := make(map[int]int)
	for _, v := range guard.sleepLog {
		for m := v.start; m < v.stop; m++ {
			if val, ok := minuteMap[m]; ok {
				minuteMap[m] = val + 1
			} else {
				minuteMap[m] = 1
			}
		}
	}

	largestMinute := -1
	max := -1
	for k, v := range minuteMap {
		if v > max {
			max = v
			largestMinute = k
		}
	}
	return largestMinute, max //142515
}
