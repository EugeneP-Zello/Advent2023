package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Hand struct {
	cards []int
	win   int
	rank  int
}

func (h *Hand) calcRank(part1 bool) int {
	cards := make(map[int]int)
	for _, c := range h.cards {
		cards[c]++
	}
	rank := 0
	if !part1 {
		r0 := cards[1]
		if r0 > 0 {
			delete(cards, 1)
		}
		max_index := 0
		vmax := 0
		for idx, v := range cards {
			if v*v > vmax {
				max_index = idx
				vmax = v * v
			}
		}
		cards[max_index] += r0
	}

	for _, v := range cards {
		rank += (v * v)
	}

	return rank
}

func (h *Hand) less(other *Hand) bool {
	if h.rank > other.rank {
		return true
	} else if h.rank < other.rank {
		return false
	}
	for i := 0; i < 5; i++ {
		if h.cards[i] > other.cards[i] {
			return true
		} else if h.cards[i] < other.cards[i] {
			return false
		}
	}
	return false
}

func readCard(r rune, part1 bool) int {
	switch r {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		if part1 {
			return 11
		} else {
			return 1
		}
	case 'T':
		return 10
	default:
		return int(r - '0')
	}
}

func createHand(line string, part1 bool) *Hand {
	cards := make([]int, 0)
	for i := 0; i < 5; i++ {
		cards = append(cards, readCard(rune(line[i]), part1))
	}
	var win int
	_, _ = fmt.Sscanf(line[5:], "%d", &win)
	h := Hand{
		cards: cards,
		win:   win,
		rank:  0,
	}
	h.rank = h.calcRank(part1)
	return &h
}

func calc(filename string, part1 bool) int {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	hands := make([]*Hand, 0)
	for scanner.Scan() {
		s := scanner.Text()
		hands = append(hands, createHand(s, part1))
	}
	sort.Slice(hands, func(i, j int) bool {
		return !hands[i].less(hands[j])
	})
	total := 0
	for idx, h := range hands {
		total += h.win * (idx + 1)
	}
	return total
}

func main() {
	fmt.Println("day 07")
	test := calc("test.txt", true)
	fmt.Printf("\nTest result is %d\n", test)
	r1 := calc("input.txt", true)
	fmt.Printf("\nResult is %d\n", r1)
	test2 := calc("test.txt", false)
	fmt.Printf("\nTest result is %d\n", test2)
	r2 := calc("input.txt", false)
	fmt.Printf("\nResult is %d\n", r2)

}
