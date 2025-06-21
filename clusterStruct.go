package main

import "fmt"

type Cluster struct {
	Name   string
	Region string
	Nodes  int
}

func getCluster() Cluster {
	return Cluster{
		Name:   "gke-prod",
		Region: "us-central1",
		Nodes:  5,
	}
}

func main() {
	describeCluster := getCluster()
	fmt.Println("Cluster Name:", describeCluster.Name)
	fmt.Println("Region:", describeCluster.Region)
	fmt.Println("Nodes:", describeCluster.Nodes)
}
