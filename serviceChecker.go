package main

import "fmt"

func main() {
	services := []string{"nginx", "redis", "postgres"}

	for _, service := range services {
		fmt.Println("Checking status of", service, "...")
	}
}
