// azure/list_resources.go
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
		resourceCount += len(page.Value)
	}

	fmt.Printf("🔍 Total Azure Resources in Subscription %s: %d\n", subscriptionID, resourceCount)
}

