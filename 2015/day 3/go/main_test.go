package main

import (
	"testing"
)

func Test_Part1(t *testing.T) {
	if Part1("../input.txt") != 2081 {
		t.Fail()
	}
}

func Test_Part2_case1(t *testing.T) {
	if Part2("../case1.txt") != 3 {
		t.Fail()
	}
}

func Test_Part2_case2(t *testing.T) {
	if Part2("../case2.txt") != 3 {
		t.Fail()
	}
}

func Test_Part2_case3(t *testing.T) {
	if Part2("../case3.txt") != 11 {
		t.Fail()
	}
}

func Test_Part2(t *testing.T) {
	if Part2("../input.txt") != 2341 {
		t.Fail()
	}
}
