package main

import "fmt"

func checkService(service string, status string) {
	if service == "nginx" && status == "active" {
		fmt.Println("Web server is running")
	} else if service == "redis" && status == "inactive" {
		fmt.Println("redis is down")
	} else {
		fmt.Println("Unknown service or state")
	}
}

func main() {
	checkService("redis", "inactive")
}
