package main

import "fmt"

func getClusterInfo() (string, int, string) {
	name := "gke-prod"
	nodeCount := 5
	region := "us-centrall"
	return name, nodeCount, region
}
func main() {
	name, count, region := getClusterInfo()
	fmt.Println("Cluster: ", name)
	fmt.Println("Nodes: ", count)
	fmt.Println("Region: ", region)
}
