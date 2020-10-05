package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func paintNeeded(width, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)

	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}

	area := width * height
	return area / 10.0, nil
}

func main() {
	var total float64

	for x := 0; x < 3; x++ {
		fmt.Print("Width: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		width, err := strconv.ParseFloat(input, 64)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print("Height: ")
		reader = bufio.NewReader(os.Stdin)
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		height, err := strconv.ParseFloat(input, 64)
		if err != nil {
			log.Fatal(err)
		}

		amount, err := paintNeeded(width, height)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%0.2f litters needed\n\n", amount)
		total += amount

	}

	fmt.Printf("Total: %0.2f liters\n", total)
}
