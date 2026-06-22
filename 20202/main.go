package main

import (
	"bufio"
	"fmt"
	"os"
)

type OldPasswordPolicy struct {
	min      int
	max      int
	char     byte
	password string
}

func initOldPasswordPolicy(text string) OldPasswordPolicy {
	var p OldPasswordPolicy
	var charStr string
	_, err := fmt.Sscanf(text, "%d-%d %1s: %s", &p.min, &p.max, &charStr, &p.password)
	if err == nil && len(charStr) > 0 {
		p.char = charStr[0]
	}
	return p
}

func (p OldPasswordPolicy) isValid() bool {
	count := 0
	for i := 0; i < len(p.password); i++ {
		if p.password[i] == p.char {
			count++
		}
	}
	return count >= p.min && count <= p.max
}

type OfficialPasswordPolicy struct {
	firstIndex  int
	secondIndex int
	char        byte
	password    string
}

func initOfficialPasswordPolicy(text string) OfficialPasswordPolicy {
	var p OfficialPasswordPolicy
	var charStr string
	_, err := fmt.Sscanf(text, "%d-%d %1s: %s", &p.firstIndex, &p.secondIndex, &charStr, &p.password)
	if err == nil && len(charStr) > 0 {
		p.char = charStr[0]
	}
	return p
}

func (p OfficialPasswordPolicy) isValid() bool {
	firstValid := p.firstIndex-1 >= 0 && p.firstIndex-1 < len(p.password) && p.password[p.firstIndex-1] == p.char
	secondValid := p.secondIndex-1 >= 0 && p.secondIndex-1 < len(p.password) && p.password[p.secondIndex-1] == p.char
	return firstValid != secondValid
}

func part1(lines []string) {
	validPasswords := 0

	for _, line := range lines {
		p := initOldPasswordPolicy(line)
		if p.isValid() {
			validPasswords++
		}
	}

	fmt.Printf("Part 1: %d\n", validPasswords)
}

func part2(lines []string) {
	validPasswords := 0

	for _, line := range lines {
		p := initOfficialPasswordPolicy(line)
		if p.isValid() {
			validPasswords++
		}
	}

	fmt.Printf("Part 2: %d\n", validPasswords)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	part1(lines)
	part2(lines)
}
