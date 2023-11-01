package main

import (
	"aoc/utils"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(Part2("../case2.txt"))
}

func Part1(filename string) int {
	f, open_err := os.Open(filename)
	utils.Check(open_err)
	defer f.Close()

	cur := [2]int{0, 0}
	m := make(map[[2]int]bool)
	m[cur] = true
	houses := 1

	const read_size = 10
	var seek_offset int64 = 0
	var seek_err error
	var read_count int
	var read_err error

	bytes := make([]byte, read_size)

	for {
		_, seek_err = f.Seek(seek_offset, 0)
		utils.Check(seek_err)
		seek_offset += read_size

		read_count, read_err = f.Read(bytes)
		if read_err == io.EOF {
			break
		}
		utils.Check(read_err)

	ReadByteLoop:
		for i, b := range bytes {
			if i >= read_count {
				break
			}

			switch b {
			case '^':
				cur[1]++
			case '>':
				cur[0]++
			case 'v':
				cur[1]--
			case '<':
				cur[0]--
			default:
				break ReadByteLoop
			}

			if _, hasKey := m[cur]; !hasKey {
				m[cur] = true
				houses++
			}
		}

	}

	return houses
}

func Part2(filename string) int {
	f, open_err := os.Open(filename)
	utils.Check(open_err)
	defer f.Close()

	santa := [2]int{0, 0}
	robo := [2]int{0, 0}
	isSantasTurn := true
	runner := &santa
	m := make(map[[2]int]bool)
	m[*runner] = true
	houses := 1

	const read_size = 11
	var seek_offset int64 = 0
	var seek_err error
	var read_count int
	var read_err error

	bytes := make([]byte, read_size)

	var toggle = func() {
		isSantasTurn = !isSantasTurn
	}

	for {
		_, seek_err = f.Seek(seek_offset, 0)
		utils.Check(seek_err)
		seek_offset += read_size

		// `Read` doesn't clear the byte array, so if it hasn't written new values into the whole array,
		// values from the last read was still in there. That's when `read_count` is needed.
		read_count, read_err = f.Read(bytes)
		if read_err == io.EOF {
			break
		}
		utils.Check(read_err)

	ReadByteLoop:
		for i, b := range bytes {
			if i >= read_count {
				break
			}

			if isSantasTurn {
				runner = &santa
			} else {
				runner = &robo
			}

			switch b {
			case '^':
				(*runner)[1]++
				toggle()
			case '>':
				(*runner)[0]++
				toggle()
			case 'v':
				(*runner)[1]--
				toggle()
			case '<':
				(*runner)[0]--
				toggle()
			default:
				break ReadByteLoop
			}

			if _, hasKey := m[*runner]; !hasKey {
				m[*runner] = true
				houses++
			}
		}
	}

	return houses
}
