package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getTotal(numbers []string) int {
	total := 0
	for i := 0; i < len(numbers); i++ {
		num, err := strconv.Atoi(numbers[i])
		check(err)
		total += num
	}
	return total
}

func main() {
	dat, err := ioutil.ReadFile("day1.dat")
	check(err)
	numbers := strings.Split(string(dat), "\r\n")
	total := getTotal(numbers)
	fmt.Printf("Total = %d\n", total)

	// find repeated numbers
	curFreq := 0
	freqList := []int{0}
	for {
		for i := 0; i < len(numbers); i++ {
			num, err := strconv.Atoi(numbers[i])
			check(err)
			curFreq += num
			if contains(freqList, curFreq) {
				fmt.Printf("Calibrated Frequency = %d\n", curFreq)
				os.Exit(0)
			}
			freqList = append(freqList, curFreq)
		}
	}
}

func contains(slice []int, val int) bool {
	for _, num := range slice {
		if num == val {
			return true
		}
	}
	return false
}
