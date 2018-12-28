package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// | - / \ + ^ v < >

type trackType int

type cartDirection int

type turnDirection int

const (
	NONE            trackType = iota
	VERTICAL                  // |
	HORIZONTAL                // -
	CLOCKWISE_CURVE           // /
	COUNTER_CURVE             // \
	INTERSECTION              //  +
)

const (
	NORTH cartDirection = iota // ^
	EAST                       // >
	SOUTH                      // v
	WEST                       // <
)

const (
	LEFT turnDirection = iota
	STRAIGHT
	RIGHT
)

type gridLoc struct {
	x, y int
}

type trackPart struct {
	xy gridLoc
	t  trackType
}

type track struct {
	tracks     map[gridLoc]trackPart
	carts      cartList
	maxX, maxY int
}

type cart struct {
	id       int
	xy       gridLoc
	dir      cartDirection
	nextTurn turnDirection
	collided bool
}

type cartList []cart

const interactive = false

func main() {
	dat, err := ioutil.ReadFile("day13.dat")
	check(err)
	lines := strings.Split(string(dat), "\r\n")
	part2(lines)
}

func part1(lines []string) {
	t, err := readTrack(lines)
	check(err)
	ticks := 0
	if interactive {
		fmt.Printf("Tick: %d\n", ticks)
		fmt.Println(t.String())
	}
	reader := bufio.NewReader(os.Stdin)
	done := false
	for !done {
		ticks++
		col, exists := t.runTick()
		if exists {
			done = true
			// My coords are backwards from theirs
			fmt.Println(t.String())
			fmt.Printf("Collision on tick %d at %d, %d\n", ticks, col.y, col.x)
			continue
		}

		if interactive {
			fmt.Printf("Tick: %d\n", ticks)
			fmt.Println(t.String())
		}
		//col, exists := t.carts.hasCollision()

		if interactive {
			text, _ := reader.ReadString('\n')
			if len(text) > 2 {
				done = true
			}
		}
	}
}

func part2(lines []string) {
	t, err := readTrack(lines)
	check(err)
	ticks := 0
	if interactive {
		fmt.Printf("Tick: %d\n", ticks)
		fmt.Println(t.String())
	}
	reader := bufio.NewReader(os.Stdin)
	done := false
	for !done {
		ticks++
		t.runTickRemoveCollisions()
		if len(t.carts) == 1 {
			fmt.Printf("Last cart on tick %d at %d,%d\n", ticks, t.carts[0].xy.y, t.carts[0].xy.x)
			return
		}

		if interactive {
			fmt.Printf("Tick: %d\n", ticks)
			fmt.Println(t.String())
		}

		if interactive {
			text, _ := reader.ReadString('\n')
			if len(text) > 2 {
				done = true
			}
		}
	}
}

func readTrack(lines []string) (track, error) {
	var track track
	track.tracks = make(map[gridLoc]trackPart)
	track.maxX = len(lines)
	track.maxY = len(lines[0])
	for x := range lines {
		for y := range lines[x] {
			if len(lines[x]) > track.maxY {
				track.maxY = len(lines[x])
			}
			part, err := readTrackPart(string(lines[x][y]), x, y)
			if err != nil {
				return track, err
			}
			if part.t != NONE {
				track.tracks[part.xy] = part
			}
			cart, ok := readCart(string(lines[x][y]), x, y)
			if ok {
				track.carts = append(track.carts, cart)
			}
		}
	}
	sort.Sort(track.carts)
	return track, nil
}

func readTrackPart(c string, x, y int) (trackPart, error) {
	tp := trackPart{xy: gridLoc{x: x, y: y}}

	switch c {
	case "|", "^", "v":
		tp.t = VERTICAL
	case "-", "<", ">":
		tp.t = HORIZONTAL
	case "/":
		tp.t = CLOCKWISE_CURVE
	case "\\":
		tp.t = COUNTER_CURVE
	case "+":
		tp.t = INTERSECTION
	case " ":
		tp.t = NONE
	default:
		return tp, fmt.Errorf("Unexpected character '%v' at location %d, %d", c, x, y)
	}
	return tp, nil
}

func readCart(c string, x, y int) (cart, bool) {
	id := rand.Intn(100000)
	cart := cart{id: id, xy: gridLoc{x: x, y: y}}
	found := false
	switch c {
	case "^":
		cart.dir = NORTH
		found = true
	case ">":
		cart.dir = EAST
		found = true
	case "v":
		cart.dir = SOUTH
		found = true
	case "<":
		cart.dir = WEST
		found = true
	}
	return cart, found
}

func (tp trackPart) String() string {
	switch tp.t {
	case VERTICAL:
		return "|"
	case HORIZONTAL:
		return "-"
	case CLOCKWISE_CURVE:
		return "/"
	case COUNTER_CURVE:
		return "\\"
	case INTERSECTION:
		return "+"
	case NONE:
		return " "
	}
	return " "
}

func (c cart) String() string {
	switch c.dir {
	case NORTH:
		return "^"
	case EAST:
		return ">"
	case SOUTH:
		return "v"
	case WEST:
		return "<"
	}
	return ""
}

func (g gridLoc) Less(g2 gridLoc) bool {
	if g.y == g2.y {
		return g.x < g2.x
	}
	return g.y < g2.y
}

func (c cartList) Len() int {
	return len(c)
}

func (c cartList) Less(i, j int) bool {
	return c[i].xy.Less(c[j].xy)
}

func (c cartList) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c cartList) RemoveCollisions() cartList {
	i := 0
	l := len(c)
	for i < l {
		if c[i].collided {
			c = append(c[:i], c[i+1:]...)
			l--
		} else {
			i++
		}
	}
	return c[:i]
}

func (c cartList) GetAtLocation(loc gridLoc) (*cart, bool) {
	for _, v := range c {
		if v.xy == loc {
			return &v, true
		} else if !v.xy.Less(loc) {
			return nil, false
		}
	}
	return nil, false
}

func (c cartList) hasCollision() (gridLoc, bool) {
	for _, v := range c {
		carts := c.GetAllAtLocation(v.xy)
		if len(carts) > 1 {
			return v.xy, true
		}
	}
	return gridLoc{x: -1, y: -1}, false
}

func (c cartList) GetAllAtLocation(loc gridLoc) []cart {
	var carts []cart
	for _, v := range c {
		if v.xy == loc {
			carts = append(carts, v)
		}
	}
	return carts
}

func (t track) String() string {
	var out strings.Builder
	for i := 0; i < t.maxX; i++ {
		for j := 0; j < t.maxY; j++ {
			// Draw the cart first
			carts := t.carts.GetAllAtLocation(gridLoc{x: i, y: j})
			if len(carts) > 0 {
				if len(carts) > 1 {
					fmt.Fprintf(&out, "X")
				} else {
					fmt.Fprintf(&out, carts[0].String())
				}
			} else {
				val, ok := t.tracks[gridLoc{x: i, y: j}]
				if ok {
					fmt.Fprintf(&out, val.String())
				} else {
					fmt.Fprintf(&out, " ")
				}
			}
		}
		fmt.Fprintf(&out, "\n")
	}
	return out.String()
}

func (t track) runTick() (gridLoc, bool) {
	for i := range t.carts {
		t.carts[i].runTick(t.tracks[t.carts[i].xy])
		loc, collision := t.carts.hasCollision()
		if collision {
			return loc, collision
		}
	}
	// resort them after the moves
	sort.Sort(t.carts)
	return gridLoc{x: -1, y: -1}, false
}

func (t *track) runTickRemoveCollisions() {
	for i := range t.carts {
		t.carts[i].runTick(t.tracks[t.carts[i].xy])
		cartsAtLoc := t.carts.GetAllAtLocation(t.carts[i].xy)

		collidedAtLoc := false
		for j := range cartsAtLoc {
			// See if a collision actually happened, there may be other carts at this location
			// that already collided and then got cleaned up, but haven't been removed from the
			// list yet
			if cartsAtLoc[j].id != t.carts[i].id && !cartsAtLoc[j].collided {
				collidedAtLoc = true
				break
			}
		}
		if collidedAtLoc {
			t.markCollision(cartsAtLoc)
		}
	}
	// Remove collided carts
	t.carts = t.carts.RemoveCollisions() // by reference?

	// resort them after the moves
	sort.Sort(t.carts)
}

func (t track) markCollision(collidedCarts []cart) {
	for j := range collidedCarts {
		for i := range t.carts {
			if collidedCarts[j].id == t.carts[i].id {
				t.carts[i].markCollision()
			}
		}
	}
}

func (c *cart) markCollision() {
	c.collided = true
}

func (c *cart) runTick(tp trackPart) {
	switch tp.t {
	case VERTICAL:
		if c.dir == NORTH || c.dir == SOUTH {
			c.moveDir()
		} else {
			log.Fatalf("Wrong direction %v on track %v", c.dir, tp.t)
		}
	case HORIZONTAL:
		if c.dir == EAST || c.dir == WEST {
			c.moveDir()
		} else {
			log.Fatalf("Wrong direction %v on track %v", c.dir, tp.t)
		}
	case CLOCKWISE_CURVE:
		c.dir = turnClockwise(c.dir)
		c.moveDir()
	case COUNTER_CURVE:
		c.dir = turnCounterClockwise(c.dir)
		c.moveDir()
	case INTERSECTION:
		c.dir = turnIntersection(c.nextTurn, c.dir)
		c.moveDir()
		c.nextTurn = getNextTurn(c.nextTurn)
	}
}

func (c *cart) moveDir() {
	if c.dir == NORTH {
		c.xy.x = c.xy.x - 1
	} else if c.dir == SOUTH {
		c.xy.x = c.xy.x + 1
	} else if c.dir == EAST {
		c.xy.y = c.xy.y + 1
	} else if c.dir == WEST {
		c.xy.y = c.xy.y - 1
	}
}

func getNextTurn(t turnDirection) turnDirection {
	if t == RIGHT {
		return LEFT
	}
	return t + 1
}

func turnIntersection(t turnDirection, dir cartDirection) cartDirection {
	switch t {
	case LEFT:
		return turnLeft(dir)
	case STRAIGHT:
		return dir
	case RIGHT:
		return turnRight(dir)
	}
	log.Fatalf("NO valid turn direct %v", t)
	return dir
}

func turnRight(dir cartDirection) cartDirection {
	if dir == WEST {
		return NORTH
	}
	return dir + 1
}

func turnLeft(dir cartDirection) cartDirection {
	if dir == NORTH {
		return WEST
	}

	return dir - 1
}

//   /
func turnClockwise(dir cartDirection) cartDirection {
	switch dir {
	case NORTH:
		return EAST
	case EAST:
		return NORTH
	case SOUTH:
		return WEST
	case WEST:
		return SOUTH
	}
	return dir
}

//  \
func turnCounterClockwise(dir cartDirection) cartDirection {
	switch dir {
	case NORTH:
		return WEST
	case EAST:
		return SOUTH
	case SOUTH:
		return EAST
	case WEST:
		return NORTH
	}
	return dir
}
