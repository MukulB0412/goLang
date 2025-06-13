package main

import "fmt"

func checkDiskSpace(used int, total int) {
	percentage := (used * 100) / total
	if percentage > 95 {
		fmt.Println("Disk Critical")
	} else if percentage >= 70 && percentage <= 95 {
		fmt.Println("Disk usage high")
	} else if percentage < 70 {
		fmt.Println("Disk Healthy")
	}
}

func main() {
	checkDiskSpace(180, 200)
}
