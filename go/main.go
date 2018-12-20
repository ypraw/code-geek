package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("\t\tProgramming Test Interview")
	fmt.Println("\t\t\tLanguage: Go")
	hr()

	measure(sort_in_one_loop)
	hr()
}

// hr function use to print dividing line
func hr() {
	fmt.Printf("============================================================\n")
}

// measure function use to measure elapsed time to run a function
func measure(f func()) {
	start_time := time.Now()
	f()
	elapsed := float64(time.Since(start_time).Nanoseconds()) / 1000000
	fmt.Println("elapsed time: ", elapsed, "ms")
}