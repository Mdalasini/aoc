package main

import (
	"bufio"
	"fmt"
	"os"
)

func traverseMatrix(matrix [][]bool, right int, down int) int {
	cols := len(matrix[0])
	rows := len(matrix)
	r, c := 0, 0
	trees := 0

	for r < rows {
		if matrix[r][c] {
			trees++
		}

		r += down
		c = (c + right) % cols
	}

	return trees
}

func part1(matrix [][]bool) {
	trees := traverseMatrix(matrix, 3, 1)
	fmt.Printf("Part 1: %d\n", trees)
}

func part2(matrix [][]bool) {
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	result := 1
	for _, slope := range slopes {
		result *= traverseMatrix(matrix, slope[0], slope[1])
	}

	fmt.Printf("Part 2: %d\n", result)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	matrix := [][]bool{}

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]bool, len(line))
		for i, char := range line {
			if char == '#' {
				row[i] = true
			}
		}
		matrix = append(matrix, row)
	}

	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, scanner.Err())
		os.Exit(1)
	}

	part1(matrix)
	part2(matrix)
}
