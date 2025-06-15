package main

import "fmt"

func main() {
	ports := []int{22, 80, 443, 8080}

	for _, port := range ports {
		fmt.Printf("Port %d is open\n", port)
	}
}
