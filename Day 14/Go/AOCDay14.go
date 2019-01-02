package main

import (
	"fmt"
	"strings"
)

func main() {
	recipes := []int{3, 7}
	matching := []int{7, 9, 3, 0, 6, 1}
	practice := 793061

	scores := 10

	part1(practice, scores, recipes)
	part2(matching, recipes)
}

func part1(practice, scores int, recipes []int) {
	elfIdx1 := 0
	elfIdx2 := 1

	println(recipeString(recipes, elfIdx1, elfIdx2))

	for len(recipes) < practice+scores {
		added := newRecipes(recipes[elfIdx1], recipes[elfIdx2])
		recipes = append(recipes, added...)
		elfIdx1 = nextIndex(recipes, elfIdx1)
		elfIdx2 = nextIndex(recipes, elfIdx2)
		//println(recipeString(recipes, elfIdx1, elfIdx2))
	}

	// it is possible we added to many recipes (because we needed 1 more and we added two)
	// so quick fixup here
	recipes = recipes[:practice+scores]

	println("Part 1 solution = ", scoreString(recipes[len(recipes)-10:]))
}

func part2(matching []int, recipes []int) {
	matchingLen := len(matching)

	elfIdx1 := 0
	elfIdx2 := 1
	found := false

	for !found {
		added := newRecipes(recipes[elfIdx1], recipes[elfIdx2])

		recipes = append(recipes, added[0])
		recipeLen := len(recipes)
		if len(added) > 1 {
			if recipeLen >= matchingLen && compareSlices(matching, recipes[recipeLen-matchingLen:]) {
				found = true
				continue
			}
			recipes = append(recipes, added[1])
			recipeLen = len(recipes)

		}
		if recipeLen >= matchingLen && compareSlices(matching, recipes[recipeLen-matchingLen:]) {
			found = true
			continue
		}

		elfIdx1 = nextIndex(recipes, elfIdx1)
		elfIdx2 = nextIndex(recipes, elfIdx2)
	}

	println("Part 2 solution = ", len(recipes)-matchingLen)
}

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func scoreString(recipes []int) string {
	var s strings.Builder
	for _, r := range recipes {
		fmt.Fprintf(&s, "%d", r)
	}
	return s.String()
}

func recipeString(recipes []int, elf1 int, elf2 int) string {
	var s strings.Builder
	for i, r := range recipes {
		if i == elf1 && i == elf2 {
			fmt.Fprintf(&s, " [(%d)]", r)
		} else if i == elf1 {
			fmt.Fprintf(&s, " (%d)", r)
		} else if i == elf2 {
			fmt.Fprintf(&s, " [%d]", r)
		} else {
			fmt.Fprintf(&s, " %d", r)
		}
	}
	return s.String()
}

func newRecipes(a, b int) []int {
	c := a + b
	if c > 9 {
		return []int{1, c % 10}
	}
	return []int{c}
}

func nextIndex(recipes []int, idx int) int {
	forward := recipes[idx] + 1
	newIdx := idx + forward
	return newIdx % len(recipes)
}
