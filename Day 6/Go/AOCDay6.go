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
	dat, err := ioutil.ReadFile("day6.dat")
	check(err)
	points, min, max := parsePoints(strings.Split(string(dat), "\r\n"))

	println(len(points))
	fmt.Printf("Min = %+v\n", min)
	fmt.Printf("Max = %+v\n", max)

	min.x--
	min.y--
	max.x++
	max.y++
	sizeX := max.x - min.x
	sizeY := max.y - min.y

	g := grid{minX: min.x, minY: min.y, maxX: max.x, maxY: max.y}
	g.initGrid(sizeX, sizeY)
	g.assignPoints(points)
	println("Solution Part 1 = ", calcLargestArea(g, points))
	println("Solution Part 2 = ", part2(g, points))
	//g.printString()
	//for i := 0; i < len(points); i++ {
	//fmt.Printf("%d, %d = %q\n", points[i].x, points[i].y, points[i].id)
	//}
}

func part2(g grid, p []point) int {
	maxDist := 10000
	countOfRegions := 0

	for i := 0; i < len(g.gridLocations); i++ {
		for j := 0; j < len(g.gridLocations[i]); j++ {
			dist := sumOfAllDistances(g.minX+i, g.minY+j, p)
			if dist < maxDist {
				countOfRegions++
			}
		}
	}
	return countOfRegions
}

type point struct {
	id    rune
	x     int
	y     int
	count int
}

func getPoint(id rune, points []point) *point {
	for i := 0; i < len(points); i++ {
		if id == points[i].id {
			return &points[i]
		}
	}
	return &point{id: '⌀', count: -1}
}

func calcLargestArea(g grid, points []point) int {

	for i := 0; i < len(g.gridLocations); i++ {
		for j := 0; j < len(g.gridLocations[i]); j++ {
			p := getPoint(g.gridLocations[i][j].id, points)
			if i == 0 || j == 0 || i == len(g.gridLocations)-1 || j == len(g.gridLocations[i])-1 {
				// We are making the assumption that any location that has a value on the edge will
				// go to infinity. I think we can come up with some edge cases that don't work, but hopefully this
				// is good enough
				p.count = -1
			} else if p.count >= 0 {
				p.count++
			}
		}
	}

	largestArea := points[0].count
	for i := 1; i < len(points); i++ {
		if points[i].count > largestArea {
			largestArea = points[i].count
		}
	}
	return largestArea
}

func parsePoints(lines []string) ([]point, point, point) {
	points := []point{}
	var min, max point
	v := '⌀'
	for i := 0; i < len(lines); i++ {
		v++
		parts := strings.Split(lines[i], ", ")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		p := point{x: x, y: y, id: v, count: 0}
		points = append(points, p)
		if i == 0 {
			min.x = x
			max.x = x
			min.y = y
			max.y = y
			continue
		}
		if x < min.x {
			min.x = x
		}
		if x > max.x {
			max.x = x
		}
		if y < min.y {
			min.y = y
		}
		if y > max.y {
			max.y = y
		}
	}
	return points, min, max
}

type grid struct {
	gridLocations          [][]gridLoc
	minX, minY, maxX, maxY int
}

type gridLoc struct {
	id rune
}

func sumOfAllDistances(x int, y int, points []point) int {
	dist := taxiDistance(x, y, points[0].x, points[0].y)
	for i := 1; i < len(points); i++ {
		dist += taxiDistance(x, y, points[i].x, points[i].y)
	}
	return dist
}

func (g *grid) initGrid(sizeX int, sizeY int) {
	g.gridLocations = make([][]gridLoc, sizeX) //[sizeX][sizeY]point
	for i := 0; i < sizeX; i++ {
		g.gridLocations[i] = make([]gridLoc, sizeY)
		for j := 0; j < sizeY; j++ {
			g.gridLocations[i][j] = gridLoc{id: '⌀'}
		}
	}
}

func (g *grid) assignPoints(points []point) {
	for i := 0; i < len(g.gridLocations); i++ {
		for j := 0; j < len(g.gridLocations[i]); j++ {
			g.gridLocations[i][j].id = closestPoint(g.minX+i, g.minY+j, points)
		}
	}
}

func closestPoint(x int, y int, points []point) rune {
	minDist := taxiDistance(x, y, points[0].x, points[0].y)
	minID := points[0].id
	for i := 1; i < len(points); i++ {
		dist := taxiDistance(x, y, points[i].x, points[i].y)
		if dist == minDist {
			minID = '⌀'
		} else if dist < minDist {
			minID = points[i].id
			minDist = dist
		}
	}

	return minID
}

func taxiDistance(x1 int, y1 int, x2 int, y2 int) int {
	xDist := abs(x1 - x2)
	yDist := abs(y1 - y2)
	return xDist + yDist
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func (g grid) printString() {
	for i := 0; i < len(g.gridLocations); i++ {
		for j := 0; j < len(g.gridLocations[i]); j++ {
			fmt.Printf("%q", g.gridLocations[i][j].id)
		}
		println()
	}
}
