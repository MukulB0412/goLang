package main

import "fmt"

func main() {
	services := []string{"nginx", "redis", "postgres"}
	statuses := []string{"running", "stopped", "running"}

	for i := range services {
		fmt.Printf("%s is currently %s\n", services[i], statuses[i])
	}
}
