// Cf converts its numeric argument to Celsius and Fahrenheit
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"../tempconv"
)

func main() {
	if len(os.Args[1:]) > 0 {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			k := tempconv.Kelvin(t)
			fmt.Printf("%s = %s, %s = %s\n%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c), k, tempconv.KToC(k), k, tempconv.KToF(k))
			return
		}
	}
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		temps := strings.Split(reader.Text(), " ")
		for _, t := range temps {
			t, err := strconv.ParseFloat(t, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			k := tempconv.Kelvin(t)
			fmt.Printf("%s = %s, %s = %s\n%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c), k, tempconv.KToC(k), k, tempconv.KToF(k))
		}
	}
}
