package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func part1(nums []int, target int) {
	left := 0
	right := len(nums) - 1

	for left < right {
		currentSum := nums[left] + nums[right]

		if currentSum == target {
			multiple := nums[left] * nums[right]
			fmt.Printf("Part 1: %d\n", multiple)
			return
		} else if currentSum < target {
			left++
		} else {
			right--
		}
	}
	fmt.Println("Part 1: unable to find values adding up to 2020")
}

func part2(nums []int, target int) {
	n := len(nums)

	for i := 0; i < n-2; i++ {
		left := i + 1
		right := n - 1

		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]

			if currentSum == target {
				multiple := nums[i] * nums[left] * nums[right]
				fmt.Printf("Part 2: %d\n", multiple)
				return
			} else if currentSum < target {
				left++
			} else {
				right--
			}
		}
	}

	fmt.Println("Part 2: unable to find values adding up to 2020")
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	nums := []int{}

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		nums = append(nums, val)
	}

	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, scanner.Err())
		os.Exit(1)
	}

	sort.Ints(nums)

	part1(nums, 2020)
	part2(nums, 2020)
}
