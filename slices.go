package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Owen", "Hammond", "Alant", "Malcom"}
	values := []int{10, 20, 30, 40}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Owen"))
	fmt.Println(slices.Index(names, "Owen"))
}
