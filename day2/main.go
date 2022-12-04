package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
)

type position struct {
	horizontal int
	depth      int
}

func (p *position) up(s string) {
	u, _ := strconv.Atoi(s)
	p.depth = p.depth - u
}

func (p *position) down(s string) {
	d, _ := strconv.Atoi(s)
	p.depth = p.depth + d
}

func (p *position) forward(s string) {
	f, _ := strconv.Atoi(s)
	p.horizontal = p.horizontal + f
}

func (p *position) final() int {
	return p.horizontal * p.depth
}

func one(d []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(d))
	pos := position{}
	for scanner.Scan() {
		l := strings.Split(scanner.Text(), " ")
		switch l[0] {
		case "forward":
			pos.forward(l[1])
		case "up":
			pos.up(l[1])
		case "down":
			pos.down(l[1])
		}
	}
	return pos.final()
}

type positionTwo struct {
	horizontal int
	depth      int
	aim        int
}

func (p *positionTwo) up(s string) {
	u, _ := strconv.Atoi(s)
	p.aim = p.aim - u
}

func (p *positionTwo) down(s string) {
	d, _ := strconv.Atoi(s)
	p.aim = p.aim + d
}

func (p *positionTwo) forward(s string) {
	f, _ := strconv.Atoi(s)
	p.horizontal = p.horizontal + f
	p.depth = p.depth + (p.aim * f)
}

func (p *positionTwo) final() int {
	return p.horizontal * p.depth
}

func two(d []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(d))
	pos := positionTwo{}
	for scanner.Scan() {
		l := strings.Split(scanner.Text(), " ")
		switch l[0] {
		case "forward":
			pos.forward(l[1])
		case "up":
			pos.up(l[1])
		case "down":
			pos.down(l[1])
		}
	}
	return pos.final()
}

func main() {
	data, err := os.ReadFile("input.txt") // read input from file
	if err != nil {
		log.Fatal(err)
	}
	log.Default().Printf("Part 1: Final horizontal position multiplied by final depth: %d\n", one(data))
	log.Default().Printf("Part 2: Final horizontal position multiplied by final depth: %d\n", two(data))
}
