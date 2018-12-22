package main

import (
	"fmt"
	"strconv"
)

// part 1 values
// const players = 10
// const lastMarble = 1618

const players = 431
const lastMarble = 7095000

// this is insanely slow, but not so slow it won't complete, I think it took about 2 hours

func main() {

	var playerScores [players]int

	circle := []int{0}

	curNum := 1
	curMableIdx := 0
	for curNum <= lastMarble {
		if curNum%100000 == 0 {
			fmt.Printf("Current number = %d\r\n", curNum)
		}
		if curNum%23 == 0 {
			nextScore := curNum
			curMableIdx -= 7
			if curMableIdx < 0 {
				curMableIdx += len(circle)
			}
			nextScore += circle[curMableIdx]
			playerScores[curNum%players] += nextScore
			if curMableIdx == len(circle)-1 {
				circle = append(circle[:curMableIdx])
				curMableIdx = 0
			} else {
				circle = append(circle[:curMableIdx], circle[curMableIdx+1:]...)
			}
		} else {
			curMableIdx += 2
			if curMableIdx > len(circle) {
				curMableIdx -= len(circle)
			}
			circle = append(circle, 0)
			copy(circle[curMableIdx+1:], circle[curMableIdx:])
			circle[curMableIdx] = curNum
		}
		curNum++
	}
	fmt.Printf("Highest Score = %d\r\n", getHighestScore(playerScores))
}

func getHighestScore(scores [players]int) int {
	max := 0
	for i := range scores {
		if scores[i] > max {
			max = scores[i]
		}
	}
	return max
}

func circleString(circle []int, curIdx int) string {
	var s string
	for i := range circle {
		if i == curIdx {
			s += "(" + strconv.Itoa(circle[i]) + ")"
		} else {
			s += strconv.Itoa(circle[i])
		}
		s += " "
	}
	return s
}
