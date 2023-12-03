package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func testSymbols(ln string, pos int, len int) bool {
	for idx := pos - 1; idx <= pos+len; idx++ {
		if ln[idx] == '.' {
			continue
		} else if unicode.IsDigit(rune(ln[idx])) {
			continue
		} else {
			return true
		}
	}
	return false
}

func doLine(ln string, prev string, next string) int {
	ret := 0
	for idx := 1; idx < len(ln)-1; idx++ {
		if unicode.IsDigit(rune(ln[idx])) {
			var num int
			_, err := fmt.Sscanf(ln[idx:], "%d", &num)
			if err != nil {
				fmt.Println(err)
			}
			length := 1
			if num >= 10 {
				length = 2
				if num >= 100 {
					length = 3
				}
			}
			if testSymbols(prev, idx, length) || testSymbols(next, idx, length) || testSymbols(ln, idx, length) {
				ret += num
			}
			idx += length - 1
		}
	}
	return ret
}

func findGear(nums [][]*int, dy int, dx int, gear *int) *int {
	for y := dy - 1; y <= dy+1; y++ {
		for x := dx - 1; x <= dx+1; x++ {
			if nums[y][x] != nil && nums[y][x] != gear {
				return nums[y][x]
			}
		}
	}
	return nil
}

func calcGear(nums [][]*int, dy int, dx int) int {
	p1 := findGear(nums, dy, dx, nil)
	p2 := findGear(nums, dy, dx, p1)
	if p1 != nil && p2 != nil {
		return *p1 * *p2
	}
	return 0
}

func parseLine(ln string) []*int {
	ret := make([]*int, len(ln))
	for idx := 0; idx < len(ln); idx++ {
		if unicode.IsDigit(rune(ln[idx])) {
			num := new(int)
			_, _ = fmt.Sscanf(ln[idx:], "%d", num)
			ret[idx] = num
			if *num >= 10 {
				idx++
				ret[idx] = num
				if *num >= 100 {
					idx++
					ret[idx] = num
				}
			}
		} else {
			ret[idx] = nil
		}
	}
	return ret
}

func main() {
	fmt.Println("day 03")
	processFile("test.txt")
	processFile("input.txt")
}

func processFile(fn string) {
	file, _ := os.Open(fn)

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	total1, total2 := 0, 0
	var lines []string
	var dotLine string
	var maxx int
	for scanner.Scan() {
		s := "." + scanner.Text() + "."
		if len(lines) == 0 {
			maxx = len(s)
			dotLine = strings.Repeat(".", maxx)
			lines = append(lines, dotLine)
		}
		lines = append(lines, s)
	}
	lines = append(lines, dotLine)
	for idx := 1; idx < len(lines)-1; idx++ {
		total1 += doLine(lines[idx], lines[idx-1], lines[idx+1])
		fmt.Println(lines[idx])
	}
	var nums [][]*int
	nums = make([][]*int, len(lines))
	for ln, idx := range lines {
		nums[ln] = parseLine(idx)
	}
	for dy := 1; dy < len(lines)-1; dy++ {
		for dx := 1; dx < maxx-1; dx++ {
			if lines[dy][dx] == '*' {
				total2 += calcGear(nums, dy, dx)
			}
		}
	}
	fmt.Printf("total1 is %d, total2 is %d\n\n", total1, total2)
}
