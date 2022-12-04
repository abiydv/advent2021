package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strconv"
)

func one(d []int) int {
	count := 0
	for i := range d {
		if i+1 < len(d) {
			if d[i+1] > d[i] {
				count = count + 1
			}
		}
	}
	return count
}

func two(d []int) int {
	count := 0
	for i := range d {
		if i+3 < len(d) {
			if d[i+1]+d[i+2]+d[i+3] > d[i]+d[i+1]+d[i+2] {
				count = count + 1
			}
		}
	}
	return count
}

func main() {
	data, err := os.ReadFile("input.txt") // read input from file
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(data))
	readings := []int{}
	for scanner.Scan() {
		l, _ := strconv.Atoi(scanner.Text())
		readings = append(readings, l)
	}
	log.Default().Printf("Part 1: Measurements larger than the previous measurement: %d\n", one(readings))
	log.Default().Printf("Part 2: Sums of a three-measurement sliding window larger than the previous sum: %d\n", two(readings))
}
