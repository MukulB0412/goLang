package main

import "fmt"

func markCritical(cpu *int) {
	if *cpu > 90 {
		fmt.Println("ALERT: CPU usage critical")
		*cpu = 50 // reset it (mock)
	}
}

func main() {
	cpuUsage := 95
	markCritical(&cpuUsage)
	fmt.Println("Post Check CPU Usage:", cpuUsage)
}
