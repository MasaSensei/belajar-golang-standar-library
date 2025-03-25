package main

import (
	"fmt"
	"time"
)

func main() {
	var duration1 time.Duration = 100 * time.Second
	fmt.Println(duration1)

	var duration2 time.Duration = 10 * time.Millisecond
	fmt.Println(duration2)

	var duration3 time.Duration = duration1 - duration2
	fmt.Println(duration3)
}
