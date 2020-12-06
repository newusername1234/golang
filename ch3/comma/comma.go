package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string { // ex input: -1234567
	isSigned := false
	decimals := ""
	if decIndex := strings.Index(s, "."); decIndex >= 0 {
		decimals = s[decIndex:]
		s = s[:decIndex]
	}
	l := len(s)
	if strings.HasPrefix(s, "-") { // if signed remove its influence on length and indexing
		isSigned = true
		l--
		s = s[1:]
	}
	if l <= 3 {
		if isSigned {
			return "-" + s + decimals
		}
		return s + decimals
	}
	r := l % 3 // number of digits before first commas
	var buf bytes.Buffer
	if r > 0 { // if 0 dont write anything
		buf.WriteString(s[:r] + ",") // writes 1,
	}
	for i := r; i < l; i += 3 {
		buf.WriteString(s[i : i+3]) // write next 3 digits
		if i+3 < l {                // write a comma if not at end
			buf.WriteString(",")
		}
	}
	if isSigned {
		return "-" + buf.String() + decimals
	}
	return buf.String() + decimals
}

func main() {
	// fmt.Println(comma("123456789"))
	fmt.Println(comma2("-12345678.1234567"))
	fmt.Println(comma2("-.111"))
}
