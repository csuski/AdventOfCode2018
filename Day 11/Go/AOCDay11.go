package main

import "fmt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	part1()
	part2()
}

func part1() {
	const serial = 3463 // Puzzle input
	var g grid
	g.initPower(serial)

	maxPow := 0
	var loc xy
	for x := 1; x <= 297; x++ {
		for y := 1; y <= 297; y++ {
			pow := g.get3x3Power(x, y)
			if pow > maxPow {
				maxPow = pow
				loc.x = x
				loc.y = y
			}
		}
	}
	fmt.Printf("Location %d, %d has power of %d\r\n", loc.x, loc.y, maxPow)
}

func part2() {
	const serial = 3463 // Puzzle input
	var g grid
	g.initPower(serial)

	maxes := make(chan xy)

	for s := 1; s < 300; s++ {
		size := s
		go func() {
			fmt.Printf("Calculated size %d\r\n", size)
			maxPow := 0
			var loc xy
			for x := 1; x <= 297; x++ {
				for y := 1; y <= 297; y++ {
					pow := g.getXbyXPower(x, y, size)
					if pow > maxPow {
						maxPow = pow
						loc.x = x
						loc.y = y
						loc.size = size
						loc.pow = pow
					}
				}
			}
			maxes <- loc
		}()
	}

	var maxOfAll xy
	for i := 0; i < 299; i++ {
		val := <-maxes
		if i == 0 || maxOfAll.pow < val.pow {
			maxOfAll = val
		}
	}

	fmt.Printf("Location %d,%d,%d has power of %d\r\n",
		maxOfAll.x, maxOfAll.y, maxOfAll.size, maxOfAll.pow)
}

func getPowerLevel(x, y, serial int) int {
	rackID := x + 10
	return (((rackID*y+serial)*rackID)/100)%10 - 5
}

type grid [300][300]int

func (g *grid) initPower(serial int) {
	for i := 0; i < 300; i++ {
		for j := 0; j < 300; j++ {
			g[i][j] = getPowerLevel(i+1, j+1, serial)
		}
	}
}

type xy struct {
	x, y, size, pow int
}

func (g grid) get3x3Power(x, y int) int {
	if x+3 > 300 || y+3 > 300 {
		return -1
	}
	sum := 0

	for i := x - 1; i < x+2; i++ {
		for j := y - 1; j < y+2; j++ {
			sum += g[i][j]
		}
	}
	return sum
}

func (g grid) getXbyXPower(x, y, size int) int {
	if x+size > 300 || y+size > 300 {
		return -1
	}
	sum := 0

	for i := x - 1; i < x+size-1; i++ {
		for j := y - 1; j < y+size-1; j++ {
			sum += g[i][j]
		}
	}
	return sum
}
