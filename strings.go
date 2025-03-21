package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("eko", "ek"))
	fmt.Println(strings.Split("eko", "o"))
	fmt.Println(strings.ToLower("Eko"))
	fmt.Println(strings.ToUpper("eko"))
	fmt.Println(strings.Trim("  eko  ", " "))
	fmt.Println(strings.ReplaceAll("eko eko eko", "eko", "budy"))
}
