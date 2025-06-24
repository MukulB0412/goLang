package main

import "fmt"

func checkDiskSpace(used int, total int) {
	percentage := (used * 1000) / total
	if percentage > 950 {
		fmt.Println("Disk Critical")
	} else if percentage >= 700 && percentage <= 950 {
		fmt.Println("Disk usage high")
	} else if percentage < 700 {
		fmt.Println("Disk Healthy")
	}
}

func main() {
	checkDiskSpace(1800, 2000)
}
