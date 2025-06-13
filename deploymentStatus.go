package main

import "fmt"

func checkDeployment(status string) {
	if status == "Running" {
		fmt.Println("All Good")
	} else if status == "Pending" {
		fmt.Println("Deployment in progress")
	} else if status == "Failed" {
		fmt.Println("Deployment failde check logs.")
	} else {
		fmt.Println("Uknown Status")
	}
}

func main() {
	checkDeployment("Pending")
}
