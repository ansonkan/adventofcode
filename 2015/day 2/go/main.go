package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Part2("../input.txt"))
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func Part1(filename string) int {
	f, open_err := os.Open(filename)
	check(open_err)
	defer f.Close()

	total := 0

	scanner := bufio.NewScanner(f)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		d := make([]int, 3)
		for i, s := 0, strings.Split(scanner.Text(), "x"); i < len(s); i++ {
			if num, err := strconv.Atoi(s[i]); err == nil {
				d[i] = num
			}
		}

		slices.Sort(d)

		total += 2*(d[0]*d[1]+d[1]*d[2]+d[0]*d[2]) + d[0]*d[1]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return total
}

func Part2(filename string) int {
	f, open_err := os.Open(filename)
	check(open_err)
	defer f.Close()

	total := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		d := make([]int, 3)
		for i, s := 0, strings.Split(scanner.Text(), "x"); i < len(s); i++ {
			if num, err := strconv.Atoi(s[i]); err == nil {
				d[i] = num
			}
		}

		slices.Sort(d)

		total += 2*(d[0]+d[1]) + d[0]*d[1]*d[2]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return total
}
