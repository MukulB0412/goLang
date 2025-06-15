package main

import "fmt"

func getDiskInfo() (string, int) {
	name := "sda1"
	size := 2048
	return name, size
}

func main() {
	diskName, diskSize := getDiskInfo()
	fmt.Println("Disk Name:", diskName)
	fmt.Println("Disk size:", diskSize, "GB")
}
