package day3

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/heyrutvik/aoc2023/utils"
)

type GearRatios struct {
	Lines []string
}

func MakeGearRatios() *GearRatios {
	return &GearRatios{
		Lines: utils.ReadLines("./day3/input.txt"),
	}
}

func (g *GearRatios) Desc() {
	fmt.Println("Puzzle:  ", "Gear Ratios")
	fmt.Println("Link:    ", "https://adventofcode.com/2023/day/3")
}

func (g *GearRatios) Solve() {
	numbers := make(map[Location]int)
	for row, line := range g.Lines {
		for col, c := range line {
			g.Part2(&numbers, row, col, c)
		}
	}
	fmt.Println("Solution:", Reduce(&numbers))
}

type Location struct {
	row int
	col int
}

type Number struct {
	value    int
	location Location
}

func (n *Number) setValue(value int) {
	n.value = value
}

func (g *GearRatios) adjacentNumbers(point Location) (nums []Number) {
	locations := adjacentLocations(point)
	for _, loc := range locations {
		line := g.Lines[loc.row]
		num, col, err := numberAt(line, loc.col)
		if err == nil {
			nums = append(nums, Number{num, Location{loc.row, col}})
		}
	}
	nums = utils.MakeSet[Number](nums).Elements()
	return
}

func (g *GearRatios) Part1(numbers *map[Location]int, row int, col int, symbol rune) {
	if isSymbol(symbol) {
		point := Location{row, col}
		nums := g.adjacentNumbers(point)
		for _, num := range nums {
			_, exist := (*numbers)[num.location]
			if !exist {
				(*numbers)[num.location] = num.value
			}
		}
	}
}

func (g *GearRatios) Part2(numbers *map[Location]int, row int, col int, symbol rune) {
	if isSymbol(symbol) {
		point := Location{row, col}
		nums := g.adjacentNumbers(point)

		fmt.Println("symbol", string(symbol), "before nums :", nums)

		if symbol == rune('*') && len(nums) == 2 {
			// dirty way!?
			ratio := nums[0].value * nums[1].value
			nums[0].setValue(ratio)
			nums[1].setValue(0)

			for _, num := range nums {
				_, exist := (*numbers)[num.location]
				if !exist {
					(*numbers)[num.location] = num.value
				}
			}
		}
	}
}

func Reduce(numbers *map[Location]int) int {
	total := 0
	for _, value := range *numbers {
		total = value + total
	}
	return total
}

func numberAt(str string, idx int) (num int, col int, err error) {
	max := len(str) - 1
	if idx < 0 || idx > max {
		err = fmt.Errorf("out of bound idx: %v", idx)
		return
	}

	left := utils.KeepLeftWhile(unicode.IsDigit, str[:idx])
	right := utils.KeepRightWhile(unicode.IsDigit, str[idx:]) // right starts at idx
	number := fmt.Sprintf("%s%s", left, right)

	// if the length of the `right` is empty, it means there is no number at `idx`
	// if the length of the `number` is empty, no number there either.
	if len(right) == 0 || len(number) == 0 {
		err = fmt.Errorf("number not found idx: %v", idx)
		return
	}

	col, _ = idx-len(left), idx+len(right)-1
	num, _ = strconv.Atoi(string(number))
	return
}

func isSymbol(r rune) bool {
	return r != '.' && !unicode.IsLetter(r) && !unicode.IsDigit(r)
}

func adjacentLocations(point Location) (locations []Location) {
	low, high := 0, 140
	if point.row < low || point.row > high || point.col < low || point.col > high {
		return []Location{}
	}

	rows := [3]int{max(point.row-1, low), point.row, min(point.row+1, high)}
	columns := [3]int{max(point.col-1, low), point.col, min(point.col+1, high)}
	for _, row := range rows {
		for _, col := range columns {
			locations = append(locations, Location{row, col})
		}
	}

	locations = utils.MakeSet[Location](locations).Elements()
	return
}
