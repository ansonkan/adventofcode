package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func Part1_ReadFile(filename string) int {
	dat, err := os.ReadFile(filename)
	check(err)

	str := string(dat)

	floor := 0
	for _, char := range str {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}
	}

	return floor
}

func Part1_Seek(filename string) int {
	f, open_err := os.Open(filename)
	check(open_err)
	defer f.Close()

	floor := 0
	const read_size = 100
	var seek_offset int64 = 0
	var seek_err error
	var read_err error

	b := make([]byte, read_size)

	for {
		_, seek_err = f.Seek(seek_offset, 0)
		check(seek_err)

		_, read_err = f.Read(b)
		if read_err == io.EOF {
			break
		}
		check(read_err)

		floor += findFloor(b)
		seek_offset += read_size
	}

	return floor
}

func findFloor(input []byte) int {
	floor := 0

	for _, b := range input {
		switch b {
		case '(':
			floor++
		case ')':
			floor--
		}
	}

	return floor
}

func Part2(filename string) (int64, error) {
	f, open_err := os.Open(filename)
	check(open_err)
	defer f.Close()

	floor := 0
	const read_size = 100
	var seek_offset int64 = 0
	var seek_err error
	var read_err error

	bytes := make([]byte, read_size)

	for {
		_, seek_err = f.Seek(seek_offset, 0)
		check(seek_err)

		_, read_err = f.Read(bytes)
		if read_err == io.EOF {
			break
		}
		check(read_err)

		for i, b := range bytes {
			switch b {
			case '(':
				floor++
			case ')':
				floor--
			}

			if floor < 0 {
				return seek_offset + int64(i) + 1, nil
			}
		}

		seek_offset += read_size
	}

	return 0, errors.New("has never been to the basement")
}

func main() {
	fmt.Println(Part2("../input.txt"))
}
