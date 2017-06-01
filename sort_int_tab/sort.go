package main

import (
		"fmt"
		"math/rand"
)

func main() {
	var array [10]int
	var i int = 0
	var tmp int

	for i < 10 {
		array[i] = rand.Intn(20)
		fmt.Printf("tab[%d] = %d\n", i, array[i])
		i++
	}
	i = 0;
	for i < 9 {
		if array[i] > array[i + 1] {
			tmp = array[i]
			array[i] = array[i + 1]
			array[i + 1] = tmp
			i = -1;
		}
		i++;
	}
	i = 0
	fmt.Printf("\n\n")
	for i < 10 {
		fmt.Printf("tab[%d] = %d\n", i, array[i])
		i++
	}
}