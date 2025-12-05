package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	previousResult := -1
	result := 0
	// Creates grid with extra border for easy convolution kernel
	grid := createGrid()

	for result != previousResult {
		previousResult = result
		result += remove(&grid)
	}

	fmt.Printf("Result: %d", result)
}

func createGrid() [][]bool {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := make([][]bool, 0) // the length of my input is 136
	// Top border
	grid = append(grid, make([]bool, 0))

	line := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()

		gridline := make([]bool, 0) // roundup trick from stackoverflow
		// Add side, to make convolution easy
		gridline = append(gridline, false)

		for _, char := range line {
			// set the correct bit
			gridline = append(gridline, char == '@')
		}
		// Add second side border
		gridline = append(gridline, false)

		grid = append(grid, gridline)
	}

	width := len(line) + 2
	grid = append(grid, make([]bool, width))
	grid[0] = make([]bool, width) // update top border as we now know the width

	return grid
}

func remove(grid *[][]bool) int {
	result := 0
	length := len(*grid)
	width := len((*grid)[0])
	for r := 1; r < width-1; r++ {
	inner:
		for c := 1; c < length-1; c++ {
			if !(*grid)[r][c] {
				fmt.Print(".")
				continue inner
			}

			count := 0
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if (*grid)[r+i][c+j] {
						count++
					}
				}
			}

			if count < 5 {
				(*grid)[r][c] = false
				fmt.Print("x")
				result++
			} else {
				fmt.Print("@")
			}
		}
		fmt.Println()
	}

	return result
}
