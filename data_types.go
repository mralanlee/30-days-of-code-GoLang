package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	var (
		b uint64
		f float64
		k string
	)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	// Read and save an integer, double, and String to your variables.
	b, _ = strconv.ParseUint(input[0], 10, 64)
	f, _ = strconv.ParseFloat(input[1], 64)
	k = input[2]
	// Print the sum of both integer variables on a new line.
	fmt.Println(b + i)
	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", d+f)
	// Concatenate and print the String variables on a new line
	fmt.Println(s + k)
	// The 's' variable above should be printed first.
}
