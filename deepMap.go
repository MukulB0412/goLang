package main

import "fmt"

func main() {
	services := map[string]string{
		"nginx":    "running",
		"redis":    "stopped",
		"postgres": "running",
	}
	for key, value := range services {
		fmt.Printf("%s is currently %s\n", key, value)
	}
}
