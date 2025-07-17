package azure

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

// ✅ Resource struct exported for other packages like ai.go
type Resource struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Location string `json:"location"`
}

// FetchResources returns a list of Azure resources
func FetchResources(subscriptionID string) []Resource {
	ctx := context.Background()
	var results []Resource

	cred, err := azidentity.NewAzureCLICredential(nil)
	if err != nil {
		log.Fatalf("Failed to get Azure CLI credentials: %v", err)
	}

	client, err := armresources.NewClient(subscriptionID, cred, nil)
	if err != nil {
		log.Fatalf("Failed to create Azure Resource Client: %v", err)
	}

	pager := client.NewListPager(nil)

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("Error retrieving resources: %v", err)
		}

		for _, res := range page.Value {
			results = append(results, Resource{
				Name:     safeStr(res.Name),
				Type:     safeStr(res.Type),
				Location: safeStr(res.Location),
			})
		}
	}

	return results
}

// CLI output for debugging (optional)
func ListAzureResources(subscriptionID string) {
	resources := FetchResources(subscriptionID)
	for _, r := range resources {
		fmt.Printf("🔸 Name: %-30s Type: %-45s Location: %s\n", r.Name, r.Type, r.Location)
	}
	fmt.Printf("🔍 Total Azure Resources in Subscription %s: %d\n", subscriptionID, len(resources))
}

// Helper to safely dereference *string
func safeStr(v *string) string {
	if v == nil {
		return "<nil>"
	}
	return *v
}
