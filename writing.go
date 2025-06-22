package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// 🕒 Current timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// 📝 Your log message
	logMessage := fmt.Sprintf("[%s] Service prometheus restarted due to OOM\n", timestamp)

	// ⚙️ Open file in APPEND mode, create if doesn't exist
	file, err := os.OpenFile("system.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // ✅ Close when done

	// ✍️ Write the message
	if _, err := file.WriteString(logMessage); err != nil {
		fmt.Println("Error writing log:", err)
		return
	}

	fmt.Println("Log appended to system.log")
}
