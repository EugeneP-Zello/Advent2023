package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func parseSeeds(seeds string) []int {
	s := strings.Split(strings.Split(seeds, ":")[1], " ")
	var ret []int
	for _, v := range s {
		var num int
		tmp := strings.Trim(v, " ")
		if len(tmp) > 0 {
			_, err := fmt.Sscanf(tmp, "%d", &num)
			if err != nil {
				fmt.Println("ERR" + err.Error())
			}
			ret = append(ret, num)
		}
	}
	return ret
}

type Range struct {
	srcMin, srcMax int
	dstMin, dstMax int
}

func parseRange(txt string) *Range {
	if len(txt) == 0 {
		return nil
	}
	r := Range{}
	var length int
	_, err := fmt.Sscanf(txt, "%d %d %d", &r.dstMin, &r.srcMin, &length)
	if err != nil {
		fmt.Println("ERR" + err.Error())
	}
	r.srcMax = r.srcMin + length - 1
	r.dstMax = r.dstMin + length - 1
	return &r
}

func (r *Range) process(input int) (bool, int) {
	if input >= r.srcMin && input <= r.srcMax {
		return true, r.dstMin + (input - r.srcMin)
	}
	return false, input
}

type Transform struct {
	ranges []*Range
}

func (t *Transform) process(input int) int {
	for _, r := range t.ranges {
		if ok, output := r.process(input); ok {
			return output
		}
	}
	return input
}

func parseTransform(scanner *bufio.Scanner) *Transform {
	scanner.Scan()
	name := scanner.Text()
	if len(name) == 0 {
		return nil
	}
	t := Transform{
		ranges: make([]*Range, 0),
	}
	for {
		scanner.Scan()
		txt := scanner.Text()
		if len(txt) == 0 {
			break
		}
		r := parseRange(txt)
		if r != nil {
			t.ranges = append(t.ranges, r)
		}
	}
	return &t
}

func main() {
	fmt.Println("day 05")
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	seeds := parseSeeds(scanner.Text())
	seeds2 := make([]int, len(seeds))
	copy(seeds2, seeds)
	scanner.Scan()
	transforms := make([]*Transform, 0)
	for {
		tr := parseTransform(scanner)
		if tr == nil {
			break
		}
		transforms = append(transforms, tr)
	}

	for _, tr := range transforms {
		for idx, seed := range seeds {
			seeds[idx] = tr.process(seed)
		}
	}
	fmt.Println(seeds)
	sort.Ints(seeds)
	fmt.Printf("Lowest value is %d\n\n", seeds[0])

	lowest := 4294967295
	for i := 0; i < len(seeds2); {
		min := seeds2[i]
		max := seeds2[i] + seeds2[i+1]

		for idx := min; idx < max; idx++ {
			test := idx
			for _, tr := range transforms {
				test = tr.process(test)
			}
			if test < lowest {
				lowest = test
			}
		}
		fmt.Printf("Range %d: Lowest value is %d\n", i, lowest)
		i = i + 2
	}
}
