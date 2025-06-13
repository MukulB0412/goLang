package main

import "fmt"

func main() {
	for i := 8000; i < 8006; i++ {
		fmt.Println("Scanning port", i)
	}
}
