package main

import "fmt"

type Deployment struct {
	Name     string
	Replicas int
	Version  string
}

func getDeployment() Deployment {
	return Deployment{
		Name:     "frontend",
		Replicas: 3,
		Version:  "v1.2.0",
	}
}

func main() {
	deploy := getDeployment()
	fmt.Println("App Name:", deploy.Name)
	fmt.Println("Replicas:", deploy.Replicas)
	fmt.Println("Version:", deploy.Version)
}

