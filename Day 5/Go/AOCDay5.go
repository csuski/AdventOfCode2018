package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("day5.dat")
	check(err)
	str := string(dat)
	changed := true
	for changed {
		str, changed = part1(str)
	}
	fmt.Println("Solution part 1 = ", len(str))
	start := time.Now()
	min := part2(str)
	fmt.Println("Solution part 2 = ", min, " took ", time.Since(start))
	start = time.Now()
	min = part2Concurrent(str)
	fmt.Println("Solution part 2 = ", min, " took ", time.Since(start), " concurrently")
}

func part1(data string) (string, bool) {
	var newStr string
	modified := false
	for i := 0; i < len(data); i++ {
		if i == len(data)-1 {
			newStr += string(data[i])
		} else if react(string(data[i]), string(data[i+1])) {
			i = i + 1
			modified = true
		} else {
			newStr += string(data[i])
		}
	}

	return newStr, modified
}

func react(a string, b string) bool {
	return (a != b && (a == strings.ToUpper(b) || b == strings.ToUpper(a)))
}

func part2(data string) int {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	minLength := len(data)
	for i := 0; i < len(alpha); i++ {
		newStr := data
		newStr = strings.Replace(newStr, string(alpha[i]), "", -1)
		newStr = strings.Replace(newStr, strings.ToUpper(string(alpha[i])), "", -1)

		changed := true
		for changed {
			newStr, changed = part1(newStr)
		}

		if minLength > len(newStr) {
			minLength = len(newStr)
		}
	}
	return minLength
}

// Concurrent solution isn't any faster because it is just one letter that is doing all the work :(
func part2Concurrent(data string) int {
	mins := make(chan int)

	alpha := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < len(alpha); i++ {
		letter := string(alpha[i])
		go func() {
			newStr := data
			newStr = strings.Replace(newStr, string(letter), "", -1)
			newStr = strings.Replace(newStr, strings.ToUpper(letter), "", -1)

			changed := true
			for changed {
				newStr, changed = part1(newStr)
			}
			mins <- len(newStr)
		}()
	}

	minLength := len(data)
	for i := 0; i < len(alpha); i++ {
		length := <-mins
		if minLength > length {
			minLength = length
		}
	}

	return minLength
}
