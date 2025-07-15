// main.go
package main

import (
	"fmt"
	"os"

	"cloud-resource-mapper/azure"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <azure-subscription-id>")
		os.Exit(1)
	}

	subscriptionID := os.Args[1]

	fmt.Println("📡 Scanning Azure Resources...")
	azure.ListAzureResources(subscriptionID)
}

