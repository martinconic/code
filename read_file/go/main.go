package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := 0

	for scanner.Scan() {
		_ = scanner.Text()
		lines++
	}

	fmt.Println(lines)

}
