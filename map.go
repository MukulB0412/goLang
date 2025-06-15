package main

import "fmt"

func main() {
	config := map[string]string{
		"hostname": "server-1",
		"ip":       "192.168.1.10",
		"role":     "database"}
	for name, value := range config {
		fmt.Printf("%s: %s\n", name, value)
	}
}
