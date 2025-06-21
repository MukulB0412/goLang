package main

import "fmt"

type Hardware struct {
	CPU    int
	Memory int
}

type Node struct {
	Name   string
	Region string
	Spec   Hardware
}

func getHardware() Hardware {
	return Hardware{
		CPU:    8,
		Memory: 32,
	}
}

func getNode() Node {
	return Node{
		Name:   "node-1",
		Region: "us-central1",
		Spec:   getHardware(),
	}
}

func main() {
	node := getNode()
	fmt.Println("Node Name:", node.Name)
	fmt.Println("Region:", node.Region)
	fmt.Println("CPU:", node.Spec.CPU, "cores")
	fmt.Println("Memory:", node.Spec.Memory, "GB")
}
