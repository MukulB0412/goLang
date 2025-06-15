package main

import "fmt"

func main() {
	msgs := []string{"Container started", "Healthcheck passed", "Scaling initiated"}

	for i, msg := range msgs {
		fmt.Printf("[%d] %s\n", i, msg)
	}
}
