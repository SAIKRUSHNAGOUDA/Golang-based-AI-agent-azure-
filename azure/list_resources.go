package azure

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

func ListAzureResources(subscriptionID string) {
	ctx := context.Background()

	cred, err := azidentity.NewAzureCLICredential(nil)
	if err != nil {
		log.Fatalf("Failed to get Azure CLI credentials: %v", err)
	}

	client, err := armresources.NewClient(subscriptionID, cred, nil)
	if err != nil {
		log.Fatalf("Failed to create Azure Resource Client: %v", err)
	}

	pager := client.NewListPager(nil)
	resourceCount := 0

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("Error retrieving resources: %v", err)
		}

		for _, res := range page.Value {
			// Safely dereference pointers
			name := "<no-name>"
			if res.Name != nil {
				name = *res.Name
			}
			resourceType := "<no-type>"
			if res.Type != nil {
				resourceType = *res.Type
			}
			location := "<no-location>"
			if res.Location != nil {
				location = *res.Location
			}

			fmt.Printf("🔸 Name: %-30s Type: %-45s Location: %s\n", name, resourceType, location)
			resourceCount++
		}
	}

	fmt.Printf("🔍 Total Azure Resources in Subscription %s: %d\n", subscriptionID, resourceCount)
}
