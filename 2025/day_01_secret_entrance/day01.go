package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func mod(a, b int) int {
	return (a%b + b) % b
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dial := 50
	password1 := 0
	password2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		direction := 1
		if scanner.Text()[:1] == "L" {
			direction = -1
		}

		step, err := strconv.Atoi(scanner.Text()[1:])
		if err != nil {
			log.Fatal(err)
		}

		inc2 := step / 100
		step = step % 100

		prevDial := dial
		dial = dial + step*direction

		if prevDial != 0 && (dial < 0 || dial > 100) {
			inc2++
		}

		dial = mod(dial, 100)

		if dial == 0 && prevDial != 0 {
			password1++
			inc2++
		}

		fmt.Printf("dial : %d, incremented %d\n", dial, inc2)
		password2 += inc2
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Password (part 1): %d\n", password1)
	fmt.Printf("Password (part 2): %d\n", password2)
}
