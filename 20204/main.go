package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func initPassport(text string) Passport {
	var p Passport
	fields := strings.FieldsSeq(text)
	for field := range fields {
		parts := strings.SplitN(field, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key, val := parts[0], parts[1]
		switch key {
		case "byr":
			p.byr = val
		case "iyr":
			p.iyr = val
		case "eyr":
			p.eyr = val
		case "hgt":
			p.hgt = val
		case "hcl":
			p.hcl = val
		case "ecl":
			p.ecl = val
		case "pid":
			p.pid = val
		case "cid":
			p.cid = val
		}
	}
	return p
}

func (p Passport) isValid() bool {
	return p.byr != "" &&
		p.iyr != "" &&
		p.eyr != "" &&
		p.hgt != "" &&
		p.hcl != "" &&
		p.ecl != "" &&
		p.pid != ""
}

func fourDigitYear(year string, min, max int) bool {
	if len(year) != 4 {
		return false
	}
	n, err := strconv.Atoi(year)
	if err != nil {
		return false
	}
	return n >= min && n <= max
}

func byrIsValid(byr string) bool {
	return fourDigitYear(byr, 1920, 2002)
}

func iyrIsValid(iyr string) bool {
	return fourDigitYear(iyr, 2010, 2020)
}

func eyrIsValid(eyr string) bool {
	return fourDigitYear(eyr, 2020, 2030)
}

func hgtIsValid(hgt string) bool {
	if strings.HasSuffix(hgt, "cm") {
		n, err := strconv.Atoi(hgt[:len(hgt)-2])
		return err == nil && n >= 150 && n <= 193
	}
	if strings.HasSuffix(hgt, "in") {
		n, err := strconv.Atoi(hgt[:len(hgt)-2])
		return err == nil && n >= 59 && n <= 76
	}
	return false
}

func hclIsValid(hcl string) bool {
	if len(hcl) != 7 || hcl[0] != '#' {
		return false
	}
	for _, c := range hcl[1:] {
		if (c < '0' || c > '9') && (c < 'a' || c > 'f') {
			return false
		}
	}
	return true
}

func eclIsValid(ecl string) bool {
	switch ecl {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		return false
	}
}

func pidIsValid(pid string) bool {
	if len(pid) != 9 {
		return false
	}
	for _, c := range pid {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func (p Passport) isValidStrict() bool {
	return p.isValid() &&
		byrIsValid(p.byr) &&
		iyrIsValid(p.iyr) &&
		eyrIsValid(p.eyr) &&
		hgtIsValid(p.hgt) &&
		hclIsValid(p.hcl) &&
		eclIsValid(p.ecl) &&
		pidIsValid(p.pid)
}

func part1(texts []string) {
	var valid int
	for _, p := range texts {
		if initPassport(p).isValid() {
			valid++
		}
	}
	fmt.Printf("Part 1: %d\n", valid)
}

func part2(texts []string) {
	var valid int
	for _, p := range texts {
		if initPassport(p).isValidStrict() {
			valid++
		}
	}
	fmt.Printf("Part 2: %d\n", valid)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var currentPassport strings.Builder
	passports := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			passports = append(passports, currentPassport.String())
			// Moving to a new passport, clear the builder
			currentPassport.Reset()
		} else {
			if currentPassport.Len() > 0 {
				currentPassport.WriteString(" ")
			}
			currentPassport.WriteString(line)
		}
	}

	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, scanner.Err())
		os.Exit(1)
	}

	// last read line
	if currentPassport.Len() > 0 {
		passports = append(passports, currentPassport.String())
		currentPassport.Reset()
	}

	part1(passports)
	part2(passports)
}
