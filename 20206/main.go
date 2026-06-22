package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1(groupAnswers [][]string) {
	sum := 0
	for _, group := range groupAnswers {
		unique := make(map[rune]bool)
		for _, answer := range group {
			for _, ch := range answer {
				unique[ch] = true
			}
		}
		sum += len(unique)
	}
	fmt.Printf("Part 1: %d\n", sum)
}

func part2(groupAnswers [][]string) {
	sum := 0
	for _, group := range groupAnswers {
		counts := make(map[rune]int)
		for _, answer := range group {
			for _, ch := range answer {
				counts[ch]++
			}
		}

		for _, count := range counts {
			if count == len(group) {
				sum++
			}
		}
	}
	fmt.Printf("Part 2: %d\n", sum)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	groupAnswers := [][]string{}
	var currentGroup []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			groupAnswers = append(groupAnswers, currentGroup)
			currentGroup = nil
		} else {
			currentGroup = append(currentGroup, line)
		}
	}

	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, scanner.Err())
		os.Exit(1)
	}

	if currentGroup != nil {
		groupAnswers = append(groupAnswers, currentGroup)
	}

	part1(groupAnswers)
	part2(groupAnswers)
}
