package	main

import "fmt"
import "os"

func main() {
	var str string
	var i int = 0

	if len(os.Args) > 1 {
		str = os.Args[1]
	} else {
		str = ""
	}
	for i < len(str) && str[i] != ' ' {
		i++
	}
	i++
	for i < len(str) {
		fmt.Printf("%c", str[i])
		i++
	}
	fmt.Printf(" ")
	i = 0
	for i < len(str) && str[i] != ' ' {
		fmt.Printf("%c", str[i])
		i++
	}
	fmt.Printf("\n")
}
