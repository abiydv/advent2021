package main

import (
	"testing"
)

func input() []int {
	return []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
}

func TestOne(t *testing.T) {
	want := 7
	got := one(input())
	if got != want {
		t.Fail()
	}
}

func TestTwo(t *testing.T) {
	want := 5
	got := two(input())
	if got != want {
		t.Fail()
	}
}
