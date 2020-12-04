package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		x, err := strconv.ParseUint(reader.Text(), 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "popcount: %v\n", err)
		}
		fmt.Println(PopCount(x))
	}
}
