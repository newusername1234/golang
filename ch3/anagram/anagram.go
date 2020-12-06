package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(a, b string) bool {
	x := strings.Split(a, "")
	sort.Strings(x)
	y := strings.Split(b, "")
	sort.Strings(y)
	return strings.Join(x, "") == strings.Join(y, "")
}

func main() {
	fmt.Println(isAnagram("heya aw", "aye hwa"))
}
