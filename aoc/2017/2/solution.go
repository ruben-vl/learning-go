package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var data = [][]int{}
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		numsRow := make([]int, 0, len(fields))
		for _, s := range fields {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			numsRow = append(numsRow, n)
		}
		data = append(data, numsRow)
	}
	fmt.Println("(Part 1) The checksum for the spreadsheet is", checksum(data))
	fmt.Println("(Part 2) The sum of evenly divisible values is", sumEvenlyDivisibleValues(data))
}

func checksum(data [][]int) int {
	diffSum := 0
	for _, row := range data {
		sort.Ints(row)
		diffSum += row[len(row)-1] - row[0]
	}
	return diffSum
}

func sumEvenlyDivisibleValues(data [][]int) int {
	divSum := 0
	for _, row := range data {
		for i1, num1 := range row {
			found := false
			for i2, num2 := range row {
				if i1 != i2 {
					if math.Mod(float64(num1), float64(num2)) == float64(0) {
						divSum += num1 / num2
						found = true
						break
					}
				}
			}
			if found {
				break
			}
		}
	}
	return divSum
}
