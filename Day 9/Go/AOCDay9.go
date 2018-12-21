package main

import (
	"fmt"
	"log"
	"strconv"
)

const players = 10
const lastMarble = 1618

func main() {

	var playerScores [players]int

	circle := []int{0}

	curNum := 1
	curMableIdx := 0
	done := false
	round := 0
	for !done {
		//fmt.Printf("Round = %d\r\n", round)
		round++
		for i := range playerScores {

			if curNum%23 == 0 {
				nextScore := curNum
				curMableIdx -= 7
				if curMableIdx < 0 {
					curMableIdx += len(circle)
				}
				nextScore += circle[curMableIdx]
				playerScores[i] += nextScore
				if curMableIdx == len(circle)-1 {
					circle = append(circle[:curMableIdx])
					curMableIdx = 0
				} else {
					circle = append(circle[:curMableIdx], circle[curMableIdx+1:]...)
				}
				fmt.Printf("Score %d for %d\r\n", nextScore, curNum)
				if nextScore == lastMarble {
					done = true
					break
				} else if nextScore > lastMarble {
					log.Fatalf("Next score (%d) is larger that score we are waiting for (%d)", nextScore, lastMarble)
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
			//fmt.Printf("[%d] %s\r\n", (i + 1), circleString(circle, curMableIdx))
		}
	}
	fmt.Printf("Solution to part 1 = %d\r\n", getHighestScore(playerScores))
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
