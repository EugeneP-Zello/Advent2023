package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Galaxy struct {
	// id is unique identifier of galaxy
	x int
	y int
}

type GalaxyMap struct {
	mp   [][]rune
	maxX int
	maxY int
}

func loadFromFile(fname string) *GalaxyMap {
	file, _ := os.Open(fname)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var mp [][]rune = make([][]rune, 0)
	rowcount := 0
	for scanner.Scan() {
		s := scanner.Text()
		if strings.Count(s, "#") == 0 {
			s = strings.ReplaceAll(s, ".", "*")
		}
		runeSlice := []rune(s)
		mp = append(mp, runeSlice)
		rowcount++
	}
	// check empty columns
	for i := 0; i < len(mp[0]); i++ {
		cnt := 0
		for ; cnt < rowcount; cnt++ {
			if mp[cnt][i] == '#' {
				break
			}
		}
		if cnt == rowcount {
			fmt.Printf("empty column %d\n", i)
			for j := 0; j < rowcount; j++ {
				mp[j][i] = '*'
			}
		}
	}
	return &GalaxyMap{
		mp,
		len(mp[0]),
		rowcount,
	}
}

func sumDistances2(galaxies []Galaxy, mp *GalaxyMap, extra int) int {
	total := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			total += findShortestPath(&galaxies[i], &galaxies[j], mp, extra)
		}
	}
	return total
}

func findShortestPath(g1 *Galaxy, g2 *Galaxy, mp *GalaxyMap, extra int) int {
	distance := 0
	col := g2.x
	if g1.x < g2.x {
		distance += calcHorDistance(g1.x, g2.x, g1.y, extra, mp)
	} else {
		col = g1.x
		distance += calcHorDistance(g2.x, g1.x, g2.y, extra, mp)
	}
	if g1.y < g2.y {
		distance += calcVertDistance(g1.y, g2.y, col, extra, mp)
	} else {
		distance += calcVertDistance(g2.y, g1.y, col, extra, mp)
	}
	//fmt.Printf("distance from %d,%d to %d,%d is %d\n", g1.x, g1.y, g2.x, g2.y, distance)
	return distance
}

func calcHorDistance(start int, end int, row int, extra int, mp *GalaxyMap) int {
	distance := 0
	if start == end {
		return 0
	}
	for i := start; i < end; i++ {
		if mp.mp[row][i] == '*' {
			distance += extra
		} else {
			distance++
		}
	}
	return distance
}

func calcVertDistance(start int, end int, col int, extra int, mp *GalaxyMap) int {
	distance := 0
	if start == end {
		return 0
	}
	for i := start; i < end; i++ {
		if mp.mp[i][col] == '*' {
			distance += extra
		} else {
			distance++
		}
	}
	return distance
}
func insertSymbol(original string, position int, symbol rune) string {
	// Convert the string to a rune slice
	runes := []rune(original)

	// Ensure the position is within bounds
	if position < 0 || position > len(runes) {
		fmt.Println("Invalid position")
		return original
	}

	// Insert the symbol at the specified position
	runes = append(runes[:position], append([]rune{symbol}, runes[position:]...)...)

	// Convert the rune slice back to a string
	result := string(runes)

	return result
}

func (g *Galaxy) distanceTo(g2 *Galaxy) int {
	return int(math.Abs(float64(g.x-g2.x)) + math.Abs(float64(g.y-g2.y)))
}

func loadMap(filename string) [][]rune {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var mp [][]rune = make([][]rune, 0)
	rowcount := 0
	for scanner.Scan() {
		s := scanner.Text()
		cnt := strings.Count(s, ".")
		runeSlice := []rune(s)
		mp = append(mp, runeSlice)
		rowcount++
		if cnt == len(s) {
			fmt.Printf("empty row %d\n", rowcount)
			mp = append(mp, runeSlice)
			rowcount++
		}
	}
	// check empty columns
	for i := 0; i < len(mp[0]); i++ {
		cnt := 0
		for ; cnt < rowcount; cnt++ {
			if mp[cnt][i] != '.' {
				break
			}
		}
		if cnt == rowcount {
			fmt.Printf("empty column %d\n", i)
			for j := 0; j < rowcount; j++ {
				mp[j] = append(mp[j][:i], append([]rune{'.'}, mp[j][i:]...)...)
			}
			i++
		}
	}
	return mp
}

func generateGalaxies(mp [][]rune) []Galaxy {
	ret := make([]Galaxy, 0)
	for i := 0; i < len(mp); i++ {
		for j := 0; j < len(mp[i]); j++ {
			if mp[i][j] == '#' {
				ret = append(ret, Galaxy{x: j, y: i})
			}
		}
	}
	return ret
}

func sumDistances(galaxies []Galaxy) int {
	total := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			dist := galaxies[i].distanceTo(&galaxies[j])
			//fmt.Printf("Distance from %d,%d to %d,%d is %d\n", galaxies[i].x, galaxies[i].y, galaxies[j].x, galaxies[j].y, dist)
			total += dist
		}
	}
	return total
}

func printMap(mp [][]rune) {
	for _, row := range mp {
		for _, r := range row {
			fmt.Printf("%c", r)
		}
		fmt.Println()
	}
}

func part1(filename string) {
	mp := loadMap(filename)
	printMap(mp)
	galaxies := generateGalaxies(mp)
	total := sumDistances(galaxies)
	fmt.Printf("Total distance for %s: %d\n", filename, total)

}

func part2(filename string, extra int) {
	mp := loadFromFile(filename)
	printMap(mp.mp)
	galaxies := generateGalaxies(mp.mp)
	total := sumDistances2(galaxies, mp, extra)
	fmt.Printf("Total distance for %s: %d\n", filename, total)

}

func main() {
	//part1("test.txt")
	//part2("test.txt", 2)
	part1("input.txt")
	part2("input.txt", 2)
	part2("input.txt", 1000000)
	//part1("input.txt")
}
