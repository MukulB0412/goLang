package main

import "fmt"

func reportNodeStatus(cpu int, memory int, disk int) {
	if cpu > 90 || memory > 90 || disk > 90 {
		fmt.Printf("ALERT: Resource usage high.\nCPU: %d%%, Memory: %d%%, Disk: %d%%\n", cpu, memory, disk)
	} else {
		fmt.Println("All resources are under control")
	}
}
func main() {
	reportNodeStatus(92, 60, 85)
}
