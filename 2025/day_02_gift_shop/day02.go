package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ok := scanner.Scan()
	if !ok {
		log.Fatal("empty file...")
	}

	invalid := make([]string, 0)

	ranges := strings.Split(scanner.Text(), ",")
	for _, r := range ranges {
		endpoints := strings.Split(r, "-")
		begin, err := strconv.Atoi(endpoints[0])
		if err != nil {
			log.Fatal("failed to convert")
		}

		end, err := strconv.Atoi(endpoints[1])
		if err != nil {
			log.Fatal("failed to convert")
		}

		for i := begin; i <= end; i++ {
			id := strconv.Itoa(i)
			if isInvalid(id) {
				invalid = append(invalid, id)
			}
		}
	}

	result := 0
	for _, id := range invalid {
		numId, _ := strconv.Atoi(id)
		result += numId
	}

	fmt.Printf("Result: %d\n", result)
}

func isInvalid(id string) bool {
	for l := len(id) / 2; l >= 1; l-- {
		if len(id)%l != 0 {
			continue
		}

		if checkCandidate(id, l) {
			return true
		}
	}

	return false
}

func checkCandidate(id string, l int) bool {
	for i := 1; i < len(id)/l; i++ {
		if id[:l] != id[l*i:l*(i+1)] {
			return false
		}
	}
	return true
}
