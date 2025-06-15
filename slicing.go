package main

import "fmt"

func main() {
	containers := []map[string]string{
		{"name": "nginx", "port": "80"},
		{"name": "redis", "port": "3306"},
		{"name": "postgres", "port": "5432"},
	}

	for _, container := range containers {
		fmt.Printf("Container %s runs on port %s\n", container["name"], container["port"])
	}
}
