package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x int
	y int
}

type Piping struct {
	pipes [][]rune
	start Pos
	maxX  int
	maxY  int
}

type Runner struct {
	piping *Piping
	pos    Pos
	prev   Pos
	steps  int
}

func (r *Runner) move() bool {
	if (r.prev.y+1 != r.pos.y) && r.piping.canMoveUp(r.pos) {
		r.prev = r.pos
		r.pos.y--
		r.steps++
		return true
	}
	if (r.prev.x+1 != r.pos.x) && r.piping.canMoveLeft(r.pos) {
		r.prev = r.pos
		r.pos.x--
		r.steps++
		return true
	}
	if (r.prev.y-1 != r.pos.y) && r.piping.canMoveDown(r.pos) {
		r.prev = r.pos
		r.pos.y++
		r.steps++
		return true
	}
	if (r.prev.x-1 != r.pos.x) && r.piping.canMoveRight(r.pos) {
		r.prev = r.pos
		r.pos.x++
		r.steps++
		return true
	}
	panic("No move possible")
}

func (p *Piping) canMoveUp(from Pos) bool {
	if from.y == 0 {
		return false
	}
	if p.pipes[from.y][from.x] == '|' || p.pipes[from.y][from.x] == 'L' || p.pipes[from.y][from.x] == 'J' || p.pipes[from.y][from.x] == 'S' {
		if p.pipes[from.y-1][from.x] == '|' || p.pipes[from.y-1][from.x] == '7' || p.pipes[from.y-1][from.x] == 'F' || p.pipes[from.y-1][from.x] == 'S' {
			return true
		}
	}
	return false
}

func (p *Piping) canMoveDown(from Pos) bool {
	if from.y+1 == p.maxY {
		return false
	}
	return p.canMoveUp(Pos{from.x, from.y + 1})
}

func (p *Piping) canMoveLeft(from Pos) bool {
	if from.x == 0 {
		return false
	}
	if p.pipes[from.y][from.x] == '-' || p.pipes[from.y][from.x] == '7' || p.pipes[from.y][from.x] == 'J' || p.pipes[from.y][from.x] == 'S' {
		if p.pipes[from.y][from.x-1] == '-' || p.pipes[from.y][from.x-1] == 'L' || p.pipes[from.y][from.x-1] == 'F' || p.pipes[from.y][from.x-1] == 'S' {
			return true
		}
	}
	return false
}

func (p *Piping) canMoveRight(from Pos) bool {
	if from.x+1 == p.maxX {
		return false
	}
	return p.canMoveLeft(Pos{from.x + 1, from.y})
}

func (p *Piping) print() {
	for _, row := range p.pipes {
		for _, r := range row {
			fmt.Printf("%c", r)
		}
		fmt.Println()
	}
}

func loadPiping(filename string) *Piping {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	rowCount := 0
	var start Pos = Pos{0, 0}
	var pipes [][]rune = make([][]rune, 0)
	for scanner.Scan() {
		s := scanner.Text()
		runeSlice := []rune(s)
		if x := strings.Index(s, "S"); x != -1 {
			start.x = x
			start.y = rowCount
		}
		pipes = append(pipes, runeSlice)
		rowCount++
	}
	return &Piping{pipes,
		start,
		len(pipes[0]),
		len(pipes),
	}
}

func makeDoubleGrid(piping *Piping) [][]rune {
	var dgrid [][]rune = make([][]rune, piping.maxY*2)
	for i := 0; i < piping.maxY*2; i++ {
		dgrid[i] = make([]rune, piping.maxX*2)
		if i == 0 || i == piping.maxY*2-1 {
			for j := 0; j < piping.maxX*2; j++ {
				dgrid[i][j] = '0'
			}
		} else {
			dgrid[i][0] = '0'
			dgrid[i][piping.maxX*2-1] = '0'
		}
	}
	for y := 1; y < piping.maxY*2-1; y++ {
		for x := 1; x < piping.maxX*2-1; x++ {
			dgrid[y][x] = ' '
		}
	}
	return dgrid
}

func addNearbyPoints(piping *Piping, p Pos) []Pos {
	retVal := make([]Pos, 0)
	if p.x > 0 {
		if piping.pipes[p.y][p.x-1] != '0' {
			retVal = append(retVal, Pos{p.x - 1, p.y})
		}
	}
	if p.x < piping.maxX-1 {
		if piping.pipes[p.y][p.x+1] != '0' {
			retVal = append(retVal, Pos{p.x + 1, p.y})
		}
	}
	if p.y > 0 {
		if piping.pipes[p.y-1][p.x] != '0' {
			retVal = append(retVal, Pos{p.x, p.y - 1})
		}
	}
	if p.y < piping.maxY-1 {
		if piping.pipes[p.y+1][p.x] != '0' {
			retVal = append(retVal, Pos{p.x, p.y + 1})
		}
	}
	return retVal
}

func calcEnclosedPoints(piping *Piping) int {
	var dgrid [][]rune = makeDoubleGrid(piping)
	for y := 0; y < piping.maxY; y++ {
		for x := 0; x < piping.maxX; x++ {
			if piping.pipes[y][x] != '.' {
				dgrid[y*2][x*2] = piping.pipes[y][x]
			} else if x != 0 && x != piping.maxX-1 && y != 0 && y != piping.maxY-1 {
				dgrid[y*2][x*2] = 'I'
			}
		}
	}
	for y := 0; y < piping.maxY; y++ {
		for x := 0; x < piping.maxX; x++ {
			p := Pos{x, y}
			if piping.canMoveDown(p) {
				dgrid[y*2+1][x*2] = '|'
			}
			if piping.canMoveRight(p) {
				dgrid[y*2][x*2+1] = '-'
			}
		}
	}
	show := &Piping{dgrid, Pos{0, 0}, piping.maxX * 2, piping.maxY * 2}

	fmt.Println(" flood fill")
	for y := 0; y < piping.maxY*2; y++ {
		for x := 0; x < piping.maxX*2; x++ {
			if dgrid[y][x] == '0' {
				todo := addNearbyPoints(show, Pos{x, y})
				for ; len(todo) > 0; todo = todo[1:] {
					pos := todo[0]
					r := dgrid[pos.y][pos.x]
					if r == ' ' || r == 'I' {
						fmt.Printf("filling %d %d\n", pos.x, pos.y)
						dgrid[pos.y][pos.x] = '0'
						todo = append(todo, addNearbyPoints(show, pos)...)
					}
				}
			}
		}
	}
	enclosed := 0
	for y := 0; y < piping.maxY*2; y++ {
		for x := 0; x < piping.maxX*2; x++ {
			if dgrid[y][x] == 'I' {
				enclosed++
			}
		}
	}
	show.print()
	return enclosed
}

func createClearGrid(x int, y int) [][]rune {
	var grid [][]rune = make([][]rune, y)
	for i := 0; i < y; i++ {
		grid[i] = make([]rune, x)
		for j := 0; j < x; j++ {
			grid[i][j] = '.'
		}
	}
	return grid
}

func calcStartRune(piping *Piping) rune {
	if piping.canMoveUp(piping.start) && piping.canMoveDown(piping.start) {
		return '|'
	}
	if piping.canMoveLeft(piping.start) && piping.canMoveRight(piping.start) {
		return '-'
	}
	if piping.canMoveDown(piping.start) && piping.canMoveRight(piping.start) {
		return 'F'
	}
	if piping.canMoveDown(piping.start) && piping.canMoveLeft(piping.start) {
		return '7'
	}
	if piping.canMoveUp(piping.start) && piping.canMoveRight(piping.start) {
		return 'L'
	}
	if piping.canMoveUp(piping.start) && piping.canMoveLeft(piping.start) {
		return 'J'
	}
	panic("Can't replace the start rune")
}

func distance(fn string) int {
	piping := loadPiping(fn)
	cleanPath := &Piping{
		pipes: createClearGrid(piping.maxX, piping.maxY),
		start: piping.start,
		maxX:  piping.maxX,
		maxY:  piping.maxY,
	}
	cleanPath.pipes[piping.start.y][piping.start.x] = calcStartRune(piping)
	runner1 := Runner{piping, piping.start, piping.start, 0}
	runner1.move()
	runner2 := Runner{piping, piping.start, runner1.pos, 0}
	runner2.move()
	//fmt.Printf("R1: %d,%d R2: %d,%d\n", runner1.pos.x, runner1.pos.y, runner2.pos.x, runner2.pos.y)
	for {
		cleanPath.pipes[runner1.pos.y][runner1.pos.x] = piping.pipes[runner1.pos.y][runner1.pos.x]
		cleanPath.pipes[runner2.pos.y][runner2.pos.x] = piping.pipes[runner2.pos.y][runner2.pos.x]
		runner1.move()
		runner2.move()
		if runner1.steps > 6138 {
			fmt.Printf("R1: %d,%d R2: %d,%d\n", runner1.pos.x, runner1.pos.y, runner2.pos.x, runner2.pos.y)
		}
		if runner1.pos.x == runner2.pos.x && runner1.pos.y == runner2.pos.y {
			cleanPath.pipes[runner1.pos.y][runner1.pos.x] = piping.pipes[runner1.pos.y][runner1.pos.x]
			cleanPath.print()
			x := calcEnclosedPoints(cleanPath)
			fmt.Println("Enclosed points: ", x)
			return runner1.steps //+ runner2.steps
		}
	}
}

func main() {
	//distance1 := distance("test1.txt")
	//fmt.Println("Distance 1: ", distance1)
	//distance2 := distance("test2.txt")
	//fmt.Println("Distance 2: ", distance2)
	//distance3 := distance("test3.txt")
	//fmt.Println("Distance 3: ", distance3)

	dist := distance("input.txt")
	fmt.Println("Distance P1: ", dist)

}
