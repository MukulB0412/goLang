package main

import "fmt"

func checkPort(port int) {
	if port == 80 || port == 443 {
		fmt.Println("Web traffic port")
	} else if port == 22 {
		fmt.Println("SSH port")
	} else {
		fmt.Println("Unknown port")
	}
}

func main() {
	checkPort(22)
}
