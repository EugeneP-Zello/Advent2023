package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseGame(s string) (green int, red int, blue int) {
	strs := strings.Split(s, ",")
	for _, str := range strs {
		cube := strings.Trim(str, " ")
		var clr string
		var num int
		_, _ = fmt.Sscanf(cube, "%d %s", &num, &clr)
		if clr == "green" {
			green = num
		} else if clr == "red" {
			red = num
		} else if clr == "blue" {
			blue = num
		}
	}
	return
}

func parseGames(s string) (green int, red int, blue int) {
	green, red, blue = 0, 0, 0
	strs := strings.Split(s, ";")
	for _, str := range strs {
		g, r, b := parseGame(str)
		if g > green {
			green = g
		}
		if r > red {
			red = r
		}
		if b > blue {
			blue = b
		}
	}
	return
}

func parseGameHeader(s string) (gid int, game string) {
	strs := strings.Split(s, ":")
	_, _ = fmt.Sscanf(strs[0], "Game %d", &gid)
	game = strs[1]
	return
}

func main() {
	fmt.Println("day 02")
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	total1, total2 := 0, 0
	for scanner.Scan() {
		s := scanner.Text()
		gid, game := parseGameHeader(s)
		green, red, blue := parseGames(game)
		if red <= 12 && blue <= 14 && green <= 13 {
			total1 += gid
		}
		total2 += green * red * blue
	}
	fmt.Printf("total1 is %d, total2 is %d", total1, total2)
}
