package main

import "fmt"

type Pod struct {
	Name     string
	Status   string
	Restarts int
}

func getPod() Pod {
	return Pod{
		Name:     "initialPod",
		Status:   "running",
		Restarts: 1,
	}
}

func main() {
	describePod := getPod()
	fmt.Println("Pod Name:", describePod.Name)
	fmt.Println("Status:", describePod.Status)
	fmt.Println("Restarts:", describePod.Restarts)
}
