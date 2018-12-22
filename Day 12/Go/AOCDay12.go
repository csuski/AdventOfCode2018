package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("day12.dat")
	check(err)
	lines := strings.Split(string(dat), "\r\n")

	initState := lines[0][len("initial state: "):]
	rules := make(map[string]string)

	for i := 2; i < len(lines); i++ {
		r := strings.Split(lines[i], " => ")
		if len(r) != 2 {
			log.Fatalf("Expecting 2 values in string (%v), found (%d)", lines[i], len(r))
		}
		if len(r[0]) != 5 {
			log.Fatalf("Expecting string of length 5 (%v), found (%d)", r[0], len(r[0]))
		}
		if len(r[1]) != 1 {
			log.Fatalf("Expecting string of length 1 (%v), found (%d)", r[1], len(r[1]))
		}
		rules[r[0]] = r[1]
	}
	//part1(initState, rules)
	part2(initState, rules)
}

func part1(initState string, rules map[string]string) {
	startPlantIdx := 0
	state, added := setStateWithEmpties2(initState)
	startPlantIdx -= added

	fmt.Printf("%2d: %v\r\n", 0, state)
	for i := 1; i <= 20; i++ {
		state = createNextGeneration(state, rules)
		state, added = setStateWithEmpties2(state)
		startPlantIdx -= added
		fmt.Printf("%2d: %v\r\n", i, state)
	}

	total := calcPlants(state, startPlantIdx)
	fmt.Printf("Solution to part 1 = %d", total)
}

// it is clear for part 2 that it will never complete
// so there must be a pattern. It started changing by the same
// amount at some point in time. We will detect that and then
// do a calculation to determine the result
func part2(initState string, rules map[string]string) {
	startPlantIdx := 0
	state, added := setStateWithEmpties2(initState)
	startPlantIdx -= added

	generations := 50000000000

	//fmt.Printf("%2d: %v\r\n", 0, state)
	prevVal := 0
	prevDiff := 0
	solution := 0

	for i := 1; i <= generations; i++ {
		state = createNextGeneration(state, rules)
		state, added = setStateWithEmpties2(state)
		startPlantIdx -= added
		val := calcPlants(state, startPlantIdx)
		diff := val - prevVal

		// We are assuming that if the same value is repeated it will continue
		// to repeated, which is the case with my input, but not sure if that is
		// a valid assumption for all data sets
		if diff == prevDiff {
			solution = (generations-i)*diff + val
			break
		}

		prevDiff = diff
		prevVal = val

		//fmt.Printf("Gen: %d Diff: %d Total: %d\r\n", i, val-prevVal, val)

		if i%100000 == 0 {
			fmt.Printf("On generation %d - start id %d - total %d\r\n", i, startPlantIdx, val)
			fmt.Printf("%12d: %v\r\n", i, state)
		}
	}
	fmt.Printf("Solution to part 2 = %d", solution)
}

func createNextGeneration(state string, rules map[string]string) string {
	newState := state[:2]
	for i := 0; i <= len(state)-5; i++ {
		subStr := state[i : i+5]
		a := rules[subStr]
		if a == "" {
			newState = newState + "."
		} else {
			newState = newState + a
		}
	}

	return newState + state[len(state)-2:]
}

// we always need to have three ... at the start (and end)
func setStateWithEmpties(state string) (string, int) {
	firstPlant := strings.Index(state, "#")
	numAddedToStart := 0
	if firstPlant < 3 {
		numAddedToStart = 3 - firstPlant
		state = strings.Repeat(".", numAddedToStart) + state
	}
	lastPlant := strings.LastIndex(state, "#")
	if lastPlant >= len(state)-3 {
		numToAddToEnd := lastPlant - (len(state) - 4)
		state = state + strings.Repeat(".", numToAddToEnd)
	}
	return state, numAddedToStart
}

// we always need to have exactly three ... at the start (and end)
func setStateWithEmpties2(state string) (string, int) {
	firstPlant := strings.Index(state, "#")
	state = "..." + strings.Trim(state, ".") + "..."
	newFirstPlant := strings.Index(state, "#")
	return state, newFirstPlant - firstPlant
}

func calcPlants(state string, idx int) int {
	sum := 0
	for _, v := range state {
		if string(v) == "#" {
			sum += idx
		}
		idx++
	}
	return sum
}
