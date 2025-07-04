package main

import "fmt"

func checkContainerStatus(state string, restartCount int) {
	if state == "running" {
		if restartCount > 5 {
			fmt.Println("Container running but unstable.")
		} else {
			fmt.Println("Container is healthy")
		}
	} else if state == "stopped" {
		fmt.Println("Container is stopped")
	} else {
		fmt.Println("Unknown container stat")
	}
}

func main() {
	checkContainerStatus("running", 7)
}
