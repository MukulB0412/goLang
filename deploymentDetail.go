package main

import "fmt"

func getDeploymentDetail() (string, int, string) {
	appName := "frontend"
	replicas := 3
	version := "v1.2.0"
	return appName, replicas, version
}

func main() {
	appName, replicas, version := getDeploymentDetail()
	fmt.Printf("App Name: %s\n", appName)
	fmt.Printf("Replicas: %d\n", replicas)
	fmt.Printf("version: %s", version)
}
