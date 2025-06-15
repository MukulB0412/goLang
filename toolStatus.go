package main

import "fmt"

func main() {
	tools := []string{"Docker", "Kubernetes", "Prometheus"}
	status := []string{"installed", "missing", "installed"}

	for i := range tools {
		fmt.Printf("%s is %s\n", tools[i], status[i])
	}
}
