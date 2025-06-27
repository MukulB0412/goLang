package main

import "fmt"

func getNetworkInfo() (string, int) {
	Network := "eth0"
	speed := 918
	return Network, speed
}

func main() {
	net, spd := getNetworkInfo()
	fmt.Println("Network: ", net)
	fmt.Println("Speed: ", spd, "Mbps")
}
