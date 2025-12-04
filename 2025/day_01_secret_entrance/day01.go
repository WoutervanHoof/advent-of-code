package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// TODO read file name from argument
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dial := 50
	password := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		direction := 1
		if scanner.Text()[:1] == "L" {
			direction = -1
		}

		step, err := strconv.ParseInt(scanner.Text()[1:], 10, 0)
		if err != nil {
			log.Fatal(err)
		}

		dial = (dial + int(step)*direction) % 100

		if dial == 0 {
			password++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Password: %d\n", password)
}
