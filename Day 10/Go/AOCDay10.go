package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("day10.dat")
	check(err)
	lines := strings.Split(string(dat), "\r\n")
	var points []point

	for _, line := range lines {
		points = append(points, parseLine(line))
	}

	reader := bufio.NewReader(os.Stdin)
	done := false
	for i := 1; !done; i++ {
		p := make([]point, len(points))
		copy(p, points)
		if drawGraph(p, i) {
			// Run the program and stop if we draw a picture (because we see them all line up)
			// and report the time step
			// Hit enter to keep running.
			// This isn't an automated solution, intead the human has to see what the picture is
			// It would be an interesting problem to automatically figure out the string (if you knew
			// how each letter was written it would be easier).
			// The time that is output is the answer to part 2, the string the user reads is the answer to part 1
			println("Time ", i)
			text, _ := reader.ReadString('\n')
			if len(text) > 2 {
				done = true
			}
		}
	}
}

type velocity struct {
	x, y int
}

type point struct {
	x, y int
	vel  velocity
}

func parseLine(s string) point {
	var point point
	i := strings.Index(s, "velocity")
	p := strings.Replace(s[:i][10:], ">", "", -1)
	v := strings.Replace(s[i:][10:], ">", "", -1)
	xy := strings.Split(p, ",")

	if len(xy) != 2 {
		log.Fatalf("Expecting 2 values for position in string (%v), found (%d)", p, len(xy))
	}

	x, err := strconv.Atoi(strings.Trim(xy[0], " "))
	check(err)
	y, err := strconv.Atoi(strings.Trim(xy[1], " "))
	check(err)
	point.x = x
	point.y = y

	xy = strings.Split(v, ",")
	if len(xy) != 2 {
		log.Fatalf("Expecting 2 values for position in string (%v), found (%d)", v, len(xy))
	}
	point.vel.x, err = strconv.Atoi(strings.Trim(xy[0], " "))
	check(err)
	point.vel.y, err = strconv.Atoi(strings.Trim(xy[1], " "))
	check(err)

	return point
}

func drawGraph(points []point, step int) bool {
	var minX, maxX, minY, maxY int

	for j := range points {
		points[j].x = points[j].x + points[j].vel.x*step
		points[j].y = points[j].y + points[j].vel.y*step

		if points[j].x < minX || j == 0 {
			minX = points[j].x
		}
		if points[j].x > maxX || j == 0 {
			maxX = points[j].x
		}
		if points[j].y < minY || j == 0 {
			minY = points[j].y
		}
		if points[j].y > maxY || j == 0 {
			maxY = points[j].y
		}
	}

	if isInLine(points) {

		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				if pointAtPosition(points, x, y) {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		return true
	}
	return false
}

func pointAtPosition(points []point, x, y int) bool {
	for i := range points {
		if points[i].x == x && points[i].y == y {
			return true
		}
	}
	return false
}

func isInLine(points []point) bool {
	for i := range points {
		// See if we have five points in a line including this one
		// We don't need to check negatives because the 'lower' points
		// will end going through this loop and checking all the points again
		// If we have at least five in a vertical or horizontal row then we should have a
		// letter (or the start of letters)
		if pointAtPosition(points, points[i].x+1, points[i].y) &&
			pointAtPosition(points, points[i].x+2, points[i].y) &&
			pointAtPosition(points, points[i].x+3, points[i].y) &&
			pointAtPosition(points, points[i].x+4, points[i].y) {
			return true
		}
		if pointAtPosition(points, points[i].x, points[i].y+1) &&
			pointAtPosition(points, points[i].x, points[i].y+2) &&
			pointAtPosition(points, points[i].x, points[i].y+3) &&
			pointAtPosition(points, points[i].x, points[i].y+4) {
			return true
		}
	}
	return false
}
