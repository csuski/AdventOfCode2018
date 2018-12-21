package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("day8.dat")
	check(err)
	strDigits := strings.Fields(string(dat))
	head, _ := parseNode(strDigits)
	println("Solution to part 1 = ", sumNodes([]node{head}))
	println("Solution to part 2 = ", sumValues(head))
}

type node struct {
	children []node
	metadata []int
}

func parseNode(digits []string) (node, []string) {
	if len(digits) < 2 {
		log.Fatalf("Can't parse node, not enough valued (%d)\r\n", len(digits))
	}
	kids, _ := strconv.Atoi(digits[0])
	metas, _ := strconv.Atoi(digits[1])
	var n node
	n.children = make([]node, kids)
	n.metadata = make([]int, metas)

	digits = digits[2:]

	for idx := range n.children {
		n.children[idx], digits = parseNode(digits)
	}

	for idx := range n.metadata {
		n.metadata[idx], digits = parseMeta(digits)
	}

	return n, digits
}

func parseMeta(digits []string) (int, []string) {
	if len(digits) < 1 {
		log.Fatalf("Can't parse metadata, not enough valued (%d)\r\n", len(digits))
	}
	v, _ := strconv.Atoi(digits[0])
	return v, digits[1:]
}

func sumNodes(nodes []node) int {
	sum := 0

	for i := range nodes {
		for _, v := range nodes[i].metadata {
			sum += v
		}
		sum += sumNodes(nodes[i].children)
	}
	return sum
}

func sumValues(n node) int {
	sum := 0

	if len(n.children) == 0 {
		for _, v := range n.metadata {
			sum += v
		}
	}

	for _, v := range n.metadata {
		if v <= len(n.children) {
			sum += sumValues(n.children[v-1])
		}
	}
	return sum
}
