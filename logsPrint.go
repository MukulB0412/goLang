package main

import "fmt"

func main() {
	logs := []string{"started container", "pulled image", "deployment successful"}

	for i, logs := range logs {
		fmt.Printf("Log %d: %s\n", i+1, logs)
	}
}
