package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Commands struct {
	commands string
	ptr      int
	count    int
}

func (c *Commands) next() int {
	if c.ptr >= len(c.commands) {
		c.ptr = 0
	}
	ret := c.commands[c.ptr]
	c.ptr++
	c.count++
	if ret == 'L' {
		return 0
	} else if ret == 'R' {
		return 1
	}
	panic("Invalid command")
	return -1
}

func (c *Commands) reset() {
	c.ptr = 0
	c.count = 0
}

func parseCommands(commands string) *Commands {
	return &Commands{
		commands: commands,
		ptr:      0,
		count:    0,
	}
}

type Hand struct {
	id    string
	left  string
	right string
}

func (h *Hand) IsStart1() bool {
	return h.id == "AAA"
}
func (h *Hand) IsStart2() bool {
	return h.id[2] == 'A'
}

func (h *Hand) IsFinish1() bool {
	return h.id == "ZZZ"
}

func (h *Hand) IsFinish2() bool {
	return h.id[2] == 'Z'
}

func parseHand(s string) *Hand {
	s1 := strings.Split(s, " = ")
	var id, left, right string
	id = s1[0]
	s2 := strings.Trim(s1[1], "()")
	s3 := strings.Split(s2, ", ")
	left = s3[0]
	right = s3[1]
	return &Hand{
		id:    id,
		left:  left,
		right: right,
	}
}

type State struct {
	pos   *Hand
	step1 bool
	mp    map[string]*Hand
}

func (c *State) setup(start *Hand) {
	c.pos = start
}

func (c *State) executeCommand(left bool) bool {
	if left {
		c.pos = c.mp[c.pos.left]
	} else {
		c.pos = c.mp[c.pos.right]
	}
	if c.step1 {
		return c.pos.IsFinish1()
	}
	return c.pos.IsFinish2()
}

func (c *State) print() {
	fmt.Println("Current node: " + c.pos.id)
}

func Lcm(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	lcm := Lcm(nums[1:])
	return nums[0] * lcm / gcd(nums[0], lcm)
}

func Lcm2(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func calcStepsCount(filename string, step1 bool) int {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	commands := parseCommands(scanner.Text())
	scanner.Scan() // skip empty line
	var hands map[string]*Hand = make(map[string]*Hand)
	for scanner.Scan() {
		hand := parseHand(scanner.Text())
		hands[hand.id] = hand
	}

	current := State{
		step1: step1,
		mp:    hands,
	}
	if step1 {
		current.setup(hands["AAA"])
		for {
			if current.executeCommand(commands.next() == 0) {
				return commands.count
			}
		}
	} else {
		var counts []int = make([]int, 0)
		for _, hand := range hands {
			if hand.IsStart2() {
				commands.reset()
				current.setup(hand)
				for {
					if current.executeCommand(commands.next() == 0) {
						counts = append(counts, commands.count)
						break
					}
				}
			}
		}
		for i, i2 := range counts {
			fmt.Println(i, i2)
		}
		return Lcm(counts)
	}
}

func main() {
	fmt.Println("Test steps count:", calcStepsCount("test.txt", true))
	fmt.Println("Step count:", calcStepsCount("input.txt", true))
	fmt.Println("Test2 steps count:", calcStepsCount("test2.txt", false))
	fmt.Println("Step count:", calcStepsCount("input.txt", false))
}
