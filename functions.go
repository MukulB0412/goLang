package main

import "fmt"

func checkMemory(used int, total int) {
	percentUsed := (used * 100) / total

	if percentUsed > 90 {
		fmt.Println("Memory critical")
	} else if percentUsed >= 60 && percentUsed <= 90 {
		fmt.Println("Memory Moderate")
	} else {
		fmt.Println("Memory Healthy")
	}
}
func main() {

	checkMemory(8, 16)

}
