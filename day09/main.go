package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func generateLineBelow(line []int) []int {
	ret := make([]int, len(line)-1)
	for i := 1; i < len(line); i++ {
		ret[i-1] = line[i] - line[i-1]
	}
	return ret
}

func Estimate(line []int, fwd bool) int {
	lnBelow := generateLineBelow(line)
	fmt.Println("Line below:", lnBelow)
	if lnBelow[0] == 0 && lnBelow[len(lnBelow)-1] == 0 {
		if fwd {
			return line[len(line)-1]
		}
		return line[0]
	}
	if fwd {
		return line[len(line)-1] + Estimate(lnBelow, fwd)
	}
	return line[0] - Estimate(lnBelow, fwd)
}

func calcEstimate(ln string, fwd bool) int {
	s2 := strings.Split(ln, " ")
	values := make([]int, len(s2))
	for i, s := range s2 {
		values[i], _ = strconv.Atoi(s)
	}
	return Estimate(values, fwd)
}

func calcEstimates(filename string, fwd bool) int {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	total := 0
	for scanner.Scan() {
		s := scanner.Text()
		estimate := calcEstimate(s, fwd)
		fmt.Println("Next estimate:", estimate)
		total += estimate
	}
	return total
}

func main() {
	test := calcEstimates("test.txt", true)
	fmt.Println("Test:", test)
	res := calcEstimates("input.txt", true)
	fmt.Println("P1:", res)
	test = calcEstimates("test.txt", false)
	fmt.Println("Test backward:", test)
	res = calcEstimates("input.txt", false)
	fmt.Println("P2:", res)

}
