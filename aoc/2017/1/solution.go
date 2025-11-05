package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	runes := []rune(string(b))

	sum := 0
	for i := range runes {
		if runes[i] == runes[(i+1)%len(runes)] {
			sum += int(runes[i] - '0')
		}
	}
	fmt.Println("(Part 1) Total sum:", sum)

	sum = 0
	for i := range runes {
		if runes[i] == runes[(i+len(runes)/2)%len(runes)] {
			sum += int(runes[i] - '0')
		}
	}
	fmt.Println("(Part 2) Total sum:", sum)

}
