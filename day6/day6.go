package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseArray(s string) []int {
	var ret []int
	ss := strings.Split(s, ":")[1]
	for i := 0; i < len(ss); i++ {
		if ss[i] == ' ' {
			continue
		}
		var num int
		_, _ = fmt.Sscanf(ss[i:], "%d", &num)
		if num > 9 {
			i++
			if num > 99 {
				i++
				if num > 999 {
					i++
				}
			}
		}
		ret = append(ret, num)
	}
	return ret
}

func parseArray2(s string) int {
	ss := strings.Split(s, ":")[1]
	s3 := strings.Split(ss, " ")
	var num int
	_, _ = fmt.Sscanf(strings.Join(s3, ""), "%d", &num)
	return num
}

func calcDistances(time int, dist int) int {
	count := 0
	for wait := 0; wait < time; wait++ {
		d := wait * (time - wait)
		if d > dist {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println("day 06")
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	s1 := scanner.Text()
	times := parseArray(s1)
	scanner.Scan()
	s2 := scanner.Text()
	dist := parseArray(s2)
	fmt.Println(times)
	fmt.Println(dist)
	count := 1
	for race := 0; race < len(times); race++ {
		newCount := calcDistances(times[race], dist[race])
		fmt.Printf("winning strat count: %d", newCount)
		count = count * newCount
	}
	fmt.Printf("total winning strat power: %d\n\n", count)

	time2 := parseArray2(s1)
	dist2 := parseArray2(s2)
	count2 := calcDistances(time2, dist2)
	fmt.Printf("winning strat count#2: %d\n", count2)
}
