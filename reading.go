package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1️⃣ Read the log file
	data, err := os.ReadFile("system.log")
	if err != nil {
		fmt.Println("❌ Error reading file:", err)
		return
	}

	// 2️⃣ Split file into lines
	lines := strings.Split(string(data), "\n")

	// 3️⃣ Count and print ERROR lines only
	errorCount := 1
	for _, line := range lines {
		// ✅ Check if line contains the word "ERROR"
		if strings.Contains(strings.ToLower(line), "oom") {
			fmt.Printf("[%d] %s\n", errorCount, line)
			errorCount++
		}
	}
}
