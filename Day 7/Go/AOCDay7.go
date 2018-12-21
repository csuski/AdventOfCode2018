package main

import (
	"io/ioutil"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("day7.dat")
	check(err)
	lines := strings.Split(string(dat), "\r\n")
	//part1(lines)
	part2(lines)
}

func part1(lines []string) {
	nodes := make(map[string]node)
	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")
		parent := addToTree(nodes, parts[1])
		child := addToTree(nodes, parts[7])
		addDependency(nodes, parent.letter, child.letter)
	}

	heads := getHeadsOfTree(nodes)
	sol := getLetters(nodes, "", heads)
	println("Solution part 1 = ", len(sol), " - ", (sol))

}

func part2(lines []string) {
	nodes := make(map[string]node)
	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")
		parent := addToTree(nodes, parts[1])
		child := addToTree(nodes, parts[7])
		addDependency(nodes, parent.letter, child.letter)
	}

	heads := getHeadsOfTree(nodes)
	/*sol := getLetters(nodes, "", heads)
	for _, l := range sol {
		s := string(l)
		println("Nodes = ", s, " work = ", nodes[s].workLeft)
	}*/
	time := calculateTime(nodes, heads)
	println("Solution to part 2 = ", time)
}

func calculateTime(nodes map[string]node, h []node) int {
	var workers [5]*node

	for i := 0; i < len(h) && i < 5; i++ {
		workers[i] = &h[i]
	}
	workersAvailable := 5 - len(h)
	totalWork := 0
	var lettersDone string

	for workersAvailable < 5 {
		workDone := calculateNextTime(workers)
		totalWork += workDone
		for idx := range workers {
			if workers[idx] != nil {
				workers[idx].workLeft -= workDone
				if workers[idx].workLeft == 0 {
					lettersDone += workers[idx].letter
					nodes[workers[idx].letter] = *workers[idx]
					workers[idx] = nil
					workersAvailable++
				}
			}
		}

		nextNodes := getNextNodes(nodes, lettersDone, workersAvailable, workers)
		for _, j := range nextNodes {
			for k := range workers {
				if workers[k] == nil {
					a := nodes[j]
					workers[k] = &a
					workersAvailable--
					break
				}
			}
		}
	}
	return totalWork
}

func getNextNodes(nodes map[string]node, done string, workersAvailable int, workers [5]*node) []string {
	// The next node is a node that is ready
	// Is not on the workers list already
	// has work left
	var w []string
	for k, v := range nodes {
		if v.workLeft > 0 && isReady(done, v.parents) && !onList(workers, k) {
			w = append(w, k)
		}
	}
	sort.Strings(w)
	if len(w) > workersAvailable {
		return w[:workersAvailable]
	}
	return w
}

func onList(workers [5]*node, letter string) bool {
	for i := range workers {
		if workers[i] != nil && workers[i].letter == letter {
			return true
		}
	}
	return false
}

func calculateNextTime(workers [5]*node) int {
	var t = 100
	for i := 0; i < 5; i++ {
		if workers[i] != nil && workers[i].workLeft < t {
			t = workers[i].workLeft
		}
	}
	return t
}

func isReady(done string, dependencies string) bool {
	for _, v := range dependencies {
		if !strings.Contains(done, string(v)) {
			return false
		}
	}
	return true
}

func getLetters(nodes map[string]node, done string, n []node) string {
	if n == nil || len(n) <= 0 {
		return done
	}

	idx := getNextNode(done, n)
	next := n[idx]
	n[idx] = n[len(n)-1]
	n = n[:len(n)-1]
	done += next.letter
	for i := 0; i < len(next.children); i++ {
		c := string(next.children[i])
		p := nodes[c].parents

		if isReady(done, p) {
			n = append(n, nodes[c])
		} else {

		}
	}
	return getLetters(nodes, done, n)
}

func getNextNode(done string, s []node) int {
	if len(s) == 0 {
		return -1
	}
	next := s[0]
	idx := 0
	for i, v := range s {
		if strings.Contains(done, v.letter) {
			continue
		}
		if v.letter < next.letter {
			next = v
			idx = i
		}
	}

	return idx
}

func getHeadsOfTree(nodes allNodes) []node {
	var heads []node
	for _, v := range nodes {
		if v.parents == "" {
			heads = append(heads, v)
		}
	}
	return heads
}

func addDependency(nodes allNodes, lp string, lc string) {
	p := nodes[lp]
	c := nodes[lc]
	p.children += c.letter
	nodes[lp] = p
	c.parents += p.letter
	nodes[lc] = c
}

func addToTree(nodes allNodes, letter string) *node {
	if val, ok := nodes[letter]; ok {
		return &val
	}
	n := node{letter: letter, workLeft: int(letter[0]) - 4} // A = 65, we want it to equal 61
	nodes[letter] = n
	return &n
}

type node struct {
	workLeft int
	letter   string
	parents  string // these must be done before this node
	children string // these must be done after this node
}

type allNodes map[string]node
