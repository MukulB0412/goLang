package main

import "fmt"

func main() {
	ports := map[int]string{
		22:   "open",
		80:   "open",
		443:  "closed",
		3306: "open",
	}
	for key, value := range ports {
		fmt.Printf("Port %d is %s\n", key, value)
	}
}
