package aocutils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadIntegerRows(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open %q: %w", filename, err)
	}
	defer file.Close()

	var integerRows [][]int
	scanner := bufio.NewScanner(file)
	lineNum := 1

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			lineNum++
			continue
		}

		fields := strings.Fields(line)
		integerRow := make([]int, 0, len(fields))
		for _, field := range fields {
			n, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("line %d: could not convert (%q) to integer; %w", lineNum, field, err)
			}
			integerRow = append(integerRow, n)
		}
		integerRows = append(integerRows, integerRow)
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file %q: %w", filename, err)
	}

	return integerRows, nil
}
