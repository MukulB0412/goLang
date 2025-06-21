package main

import "fmt"

type Container struct {
	Name  string
	Image string
}

type Pod struct {
	Name       string
	Containers []Container
}

func getPod() Pod {
	return Pod{
		Name: "web-app",
		Containers: []Container{
			{Name: "nginx", Image: "nginx:latest"},
			{Name: "redis", Image: "redis:alpine"},
		},
	}
}

func main() {
	pod := getPod()
	fmt.Println("Pod:", pod.Name)
	fmt.Println("Containers:")
	for _, container := range pod.Containers {
		fmt.Printf("- %s (%s)\n", container.Name, container.Image)
	}
}
