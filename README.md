# go-cloudca
Cloud.ca client for the Go programming language

# Example
```
package main

import "github.com/cloud-ca/go-cloudca"
import "github.com/cloud-ca/go-cloudca/services/cloudca"
import "fmt"

func main() {
	//Create a CcaClient
	ccaClient := gocca.NewCcaClient("[your-api-key]")
	
	//Get the available resources for a specific service and environment
	ccaResources := ccaClient.GetResources("[service-code]", "[environment-name]").(cloudca.Resources)
	
	//Get the list of instances
	instances, _ := ccaResources.Instances.List()
	fmt.Println(instances)
	
	//Get a volume with its id
	volume, _ := ccaResources.Volumes.Get("[some-volume-id]")
	fmt.Println(volume)
	
	//Create an instance
	createdInstance, _ := ccaResources.Instances.Create(cloudca.Instance{
			Name: "[new-instance-name]",
			TemplateId: "[some-template-id]",
			ComputeOfferingId:"[some-compute-offering-id]",
			NetworkId:"[some-network-id]",
		})
	fmt.Println(createdInstance)
}
```
