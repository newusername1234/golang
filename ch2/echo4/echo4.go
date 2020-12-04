// echo4 prints its command line args with custom flags
package main

import (
	"flag"
	"fmt"
	"strings"
)

var omitNewline = flag.Bool("n", false, "omit trailing newline") // -n flips omitNewline
var sep = flag.String("s", " ", "separator")                     // -s string assigns string to sep
var count = flag.Int("c", 1, "number of times")                  // -c int assigns int to count

func main() {
	flag.Parse()
	for i := 0; i < *count; i++ {
		fmt.Print(strings.Join(flag.Args(), *sep))
		if !*omitNewline {
			fmt.Println()
		}
	}
}
