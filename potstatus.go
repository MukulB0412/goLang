package main

import "fmt"

func getPodStatus() (string, string) {
	name := "nginx-Pod"
	status := "running"

	return name, status
}

func main() {
	podname, podstatus := getPodStatus()
	fmt.Println("Pod Name:", podname)
	fmt.Println("Pod status:", podstatus)
}
