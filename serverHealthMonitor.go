package main

import "fmt"

func checkServerHealth(cpu int, memory int) {
	if cpu > 85 || memory > 85 {
		fmt.Println("Server Overloaded")
	} else if cpu < 40 && memory < 40 {
		fmt.Println("Server Idle")
	} else {
		fmt.Println("Server Operating Normally.")
	}
}
func main() {
	checkServerHealth(92, 70)
}
