package main

import "os"
import "fmt"

func print_repeat(i int) {
	var len int = 1

	if os.Args[1][i] >= 'A' && os.Args[1][i] <= 'Z' {
		len = (int)(os.Args[1][i] - 'A' + 1)
	} else if os.Args[1][i] >= 'a' && os.Args[1][i] <= 'z' {
		len = (int)(os.Args[1][i] - 'a' + 1)
	}
	for len > 0 {
		fmt.Printf("%c", os.Args[1][i]);
		len--;		
	}
}

func main() {
	var i int = 0

	if len(os.Args) != 2 {
		os.Exit(0)
	}
	for i < len(os.Args[1]) {
		print_repeat(i);
		i++
	}
	fmt.Printf("\n")
}
