package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func indexOfLetter(s string) (idx int, value int) {
	idx = len(s)
	value = 0
	ss := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for v, lookup := range ss {
		i := strings.Index(s, lookup)
		if i != -1 && i < idx {
			idx = i
			value = v + 1
		}
	}
	return
}

func lastIndexOfLetter(s string) (idx int, value int) {
	idx = -1
	value = 0
	ss := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for v, lookup := range ss {
		i := strings.LastIndex(s, lookup)
		if i != -1 && i > idx {
			idx = i
			value = v + 1
		}
	}
	return
}

// it didn't work cause eightwo
// shall be 8wo on the left side and eigh2 on the right side
func replaceAll(s string) string {
	for idx := 0; idx < len(s); idx++ {
		if strings.HasPrefix(s[idx:], "one") {
			s = strings.Replace(s, "one", "1", 1)
		} else if strings.HasPrefix(s[idx:], "two") {
			s = strings.Replace(s, "two", "2", 1)
		} else if strings.HasPrefix(s[idx:], "three") {
			s = strings.Replace(s, "three", "3", 1)
		} else if strings.HasPrefix(s[idx:], "four") {
			s = strings.Replace(s, "four", "4", 1)
		} else if strings.HasPrefix(s[idx:], "five") {
			s = strings.Replace(s, "five", "5", 1)
		} else if strings.HasPrefix(s[idx:], "six") {
			s = strings.Replace(s, "six", "6", 1)
		} else if strings.HasPrefix(s[idx:], "seven") {
			s = strings.Replace(s, "seven", "7", 1)
		} else if strings.HasPrefix(s[idx:], "eight") {
			s = strings.Replace(s, "eight", "8", 1)
		} else if strings.HasPrefix(s[idx:], "nine") {
			s = strings.Replace(s, "nine", "9", 1)
		}
	}
	return s
}

func convertLine(s string) int {
	n1 := number1(s)
	n2 := number2(s)
	return n1*10 + n2
}
func number1(s string) int {
	idx := strings.IndexAny(s, "123456789")
	value, _ := strconv.Atoi(string(s[idx]))
	return value
}

func number2(s string) int {
	idx := strings.LastIndexAny(s, "123456789")
	value, _ := strconv.Atoi(string(s[idx]))
	return value
}

func convertLine2(s string) int {
	n1 := number1ex(s)
	n2 := number2ex(s)
	return n1*10 + n2
}

func number1ex(s string) int {
	idx := strings.IndexAny(s, "123456789")
	idx2, value := indexOfLetter(s)
	if idx2 > idx {
		value, _ = strconv.Atoi(string(s[idx]))
	}
	return value
}

func number2ex(s string) int {
	idx := strings.LastIndexAny(s, "123456789")
	idx2, value := lastIndexOfLetter(s)
	if idx2 < idx {
		value, _ = strconv.Atoi(string(s[idx]))
	}
	return value
}

func main() {
	fmt.Println("day 01")
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	total1, total2 := 0, 0
	for scanner.Scan() {
		s := scanner.Text()
		total1 += convertLine(s)
		total2 += convertLine2(s)
	}
	fmt.Printf("total1 is %d, total2 is %d", total1, total2)
}
