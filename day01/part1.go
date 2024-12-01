package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func parseLine(line string) (int, int) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		log.Fatalln("more than 2 number in a line")
	}
	str1, str2 := parts[0], parts[1]
	n1, err := strconv.Atoi(str1)
	if err != nil {
		log.Fatalln(err)
	}
	n2, err := strconv.Atoi(str2)
	if err != nil {
		log.Fatalln(err)
	}

	return n1, n2
}

func parseInput(input string) ([]int, []int) {
	inputReader := strings.NewReader(input)
	
	left := make([]int, 0)
	right := make([]int, 0)
	lineScanner := bufio.NewScanner(inputReader)
	for lineScanner.Scan() {
		a, b := parseLine(lineScanner.Text())
		left = append(left, a)
		right = append(right, b)
	}

	if err := lineScanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return left, right
}

func listDistance(left, right []int) int {
	// this is a destructive in-place operation, but it's OK in a throwaway program
	// sort the lists
	slices.Sort(left)
	slices.Sort(right)

	// calculate sum of distances
	var distance int
	for i := range left {
		a, b := left[i], right[i]
		var dist int
		// take distances absolute values (Go doesn't have an abs for ints)
		if a < b {
			dist = b - a
		} else {
			dist = a - b
		}
		distance += dist
	}

	return distance
}

func main() {
	fmt.Println(listDistance(parseInput()))
}
