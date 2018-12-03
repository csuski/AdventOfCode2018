package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("day3.dat")
	check(err)

	lines := strings.Split(string(dat), "\r\n")
	conflictedSquares := part1(lines)
	fmt.Println("Solution = ", conflictedSquares)
	claim := part2(lines)
	fmt.Println("Solution = ", claim)

}

func part2(lines []string) int {
	cloth := [1000][1000]int{}
	validClaims := []int{}

	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")
		claimNum, _ := strconv.Atoi(strings.TrimPrefix(parts[0], "#"))
		coord := strings.Split(parts[2], ",")
		coord[1] = strings.TrimSuffix(coord[1], ":")
		size := strings.Split(parts[3], "x")
		x, _ := strconv.Atoi(coord[0])
		y, _ := strconv.Atoi(coord[1])
		width, _ := strconv.Atoi(size[0])
		height, _ := strconv.Atoi(size[1])
		conflicted := false

		for i := x; i < x+width; i++ {
			for j := y; j < y+height; j++ {
				if cloth[i][j] == 0 {
					cloth[i][j] = claimNum
				} else {
					conflicted = true
					onList, item := containsIndex(validClaims, cloth[i][j])
					if onList {
						validClaims[item] = validClaims[len(validClaims)-1]
						validClaims = validClaims[:len(validClaims)-1]
					}
				}
			}
		}
		if !conflicted {
			validClaims = append(validClaims, claimNum)
		}

	}
	return validClaims[0]
}

func part1(lines []string) int {
	cloth := [1000][1000]int{}
	conflictList := []string{}

	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")
		coord := strings.Split(parts[2], ",")
		coord[1] = strings.TrimSuffix(coord[1], ":")
		size := strings.Split(parts[3], "x")
		x, _ := strconv.Atoi(coord[0])
		y, _ := strconv.Atoi(coord[1])
		width, _ := strconv.Atoi(size[0])
		height, _ := strconv.Atoi(size[1])

		for i := x; i < x+width; i++ {
			for j := y; j < y+height; j++ {
				if cloth[i][j] == 0 {
					cloth[i][j] = 1
				} else {
					if !contains(conflictList, strconv.Itoa(i)+"x"+strconv.Itoa(j)) {
						conflictList = append(conflictList, strconv.Itoa(i)+"x"+strconv.Itoa(j))
					}
				}
			}
		}
	}
	return len(conflictList)
}

func contains(slice []string, val string) bool {
	for _, num := range slice {
		if num == val {
			return true
		}
	}
	return false
}

func containsIndex(slice []int, val int) (bool, int) {
	for i, num := range slice {
		if num == val {
			return true, i
		}
	}
	return false, -1
}
