package main

import "fmt"

func resetIfHigh(CPU *int, Memory *int) {
	// ✅ If cpu > 90, set *cpu = 50
	if *CPU > 90 {
		*CPU = 50
	}
	if *Memory > 90 {
		*Memory = 50
	}
	// ✅ If memory > 90, set *memory = 50
}

func main() {
	cpuUsage := 94
	memUsage := 87

	resetIfHigh(&cpuUsage, &memUsage)

	fmt.Println("CPU:", cpuUsage)
	fmt.Println("Memory:", memUsage)
}
