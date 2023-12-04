package main

import (
	"bufio"
	"fmt"
	"github.com/golang-collections/collections/stack"
	"os"
	"sort"
	"strings"
)

func parseNums(nums string) []int {
	s := strings.Split(nums, " ")
	var ret []int
	for _, v := range s {
		var num int
		tmp := strings.Trim(v, " ")
		if len(tmp) > 0 {
			_, err := fmt.Sscanf(tmp, "%d", &num)
			if err != nil {
				fmt.Println(err)
			}
			ret = append(ret, num)
		}
	}
	return ret
}

func parseCard(ln string) ([]int, []int) {
	s := strings.Split(ln, "|")
	winners := parseNums(s[0])
	myNumbers := parseNums(s[1])
	sort.Ints(winners)
	sort.Ints(myNumbers)
	return winners, myNumbers
}

func lookup(nums []int, num int) bool {
	idx := sort.SearchInts(nums, num)
	return idx < len(nums) && nums[idx] == num
}

func calcScore(wins []int, my []int) (count int, score int) {
	count = 0
	score = 0
	for _, win := range wins {
		if lookup(my, win) {
			count++
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}
	return
}

func getCardNum(s string) int {
	var num int
	_, err := fmt.Sscanf(s, "Card %d", &num)
	if err != nil {
		fmt.Println(err)
	}
	return num
}

type CardStorage struct {
	cards *stack.Stack
}

func createCS() *CardStorage {
	return &CardStorage{
		cards: stack.New(),
	}
}

func (cs *CardStorage) getCardCount() int {
	tmp := cs.cards.Pop()
	if tmp == nil {
		return 0
	}
	return tmp.(int)
}

func (cs *CardStorage) wonCards(cards int, power int) {
	if cards > 0 {
		tmp := cs.cards.Pop()
		were := 0
		if tmp != nil {
			were = tmp.(int)
		}
		cs.wonCards(cards-1, power)
		cs.cards.Push(were + power)
	}
}

func main() {
	fmt.Println("day 04")
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	total1, total2 := 0, 0
	cs := createCS()
	for scanner.Scan() {
		s := scanner.Text()
		text := strings.Split(s, ":")
		card := getCardNum(text[0])
		wins, my := parseCard(text[1])
		power := cs.getCardCount() + 1
		count, score := calcScore(wins, my)
		fmt.Printf("card %d count %d score is %d\n", card, count, score)
		cs.wonCards(count, power)
		total1 += score
		total2 += power
	}
	fmt.Printf("total count is %d\n", total2)
	fmt.Printf("total score is %d\n", total1)
}
