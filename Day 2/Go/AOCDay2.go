package day2

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func boxLabelCheckSum(boxIds []string) int {
	twos := 0
	threes := 0
	for i := 0; i < len(boxIds); i++ {
		str := boxIds[i]
		addTwo := false
		addThree := false
		for j := 0; j < len(str); j++ {
			letter := str[j]
			count := 0
			for k := 0; k < len(str); k++ {
				if str[k] == letter {
					count++
				}
			}
			if count == 2 {
				addTwo = true
			}
			if count == 3 {
				addThree = true
			}
			if addTwo && addThree {
				break
			}
		}
		if addTwo {
			twos++
		}
		if addThree {
			threes++
		}
	}
	return twos * threes
}

func findMatchingLabels(boxIds []string) (string, string) {
	for i := 0; i < len(boxIds); i++ {
		for j := i + 1; j < len(boxIds); j++ {
			common := lettersInCommon(boxIds[i], boxIds[j])
			// if they are matching labels then it will return 1 less than the length
			if len(common)+1 == len(boxIds[i]) {
				return boxIds[i], boxIds[j]
			}
		}
	}

	return "", ""
}

func lettersInCommon(l1 string, l2 string) string {
	shortestLength := len(l1)
	if len(l2) < len(l1) {
		shortestLength = len(l2)
	}

	var commonLetters string
	for i := 0; i < shortestLength; i++ {
		if l1[i] == l2[i] {
			commonLetters += string(l1[i])
		}
	}
	return commonLetters
}

func main() {
	dat, err := ioutil.ReadFile("day2.dat")
	check(err)
	boxIds := strings.Split(string(dat), "\r\n")
	total := boxLabelCheckSum(boxIds)
	fmt.Printf("Solution = %d\n", total)
	label1, label2 := findMatchingLabels(boxIds)
	common := lettersInCommon(label1, label2)
	fmt.Printf("Label1 = " + label1 + " label2 = " + label2 + "\n")
	fmt.Println("Solution = " + common)

}
