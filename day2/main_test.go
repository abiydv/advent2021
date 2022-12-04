package main

import (
	"testing"
)

func input() []byte {
	return []byte("forward 5\ndown 5\nforward 8\nup 3\ndown 8\nforward 2")
}

func TestOne(t *testing.T) {
	want := 150
	got := one(input())
	if got != want {
		t.Fail()
	}
}

func TestTwo(t *testing.T) {
	want := 900
	got := two(input())
	if got != want {
		t.Fail()
	}
}
