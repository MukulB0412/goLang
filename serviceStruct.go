package main

import "fmt"

type Service struct {
	Name   string
	Ports  []int
	Status string
}

func getService() Service {
	return Service{
		Name:   "Nginx",
		Ports:  []int{80, 443},
		Status: "running",
	}
}

func main() {
	describeService := getService()
	fmt.Println("Service Name:", describeService.Name)
	fmt.Println("Ports:", describeService.Ports)
	fmt.Println("Status:", describeService.Status)
}
