# go-cloudca

Cloud.ca client for the Go programming language

# Example

```
package main

import (
	"github.com/cloud-ca/go-cloudca"
	"github.com/cloud-ca/go-cloudca/services/cloudca"
	"fmt"
	)

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

#Handling Errors
```
package main

import (
	"github.com/cloud-ca/go-cloudca"
	"github.com/cloud-ca/go-cloudca/api"
	"github.com/cloud-ca/go-cloudca/services/cloudca"
	"fmt"
	)

func main() {
	//Create a CcaClient
	ccaClient := gocca.NewCcaClient("[your-api-key]")
	
	//Get the available resources for a specific service and environment
	ccaResources := ccaClient.GetResources("[service-code]", "[environment-name]").(cloudca.Resources)

	//Get a volume with a bogus id
	_, err := ccaResources.Volumes.Get("[some-volume-id]")
	
	//Handle the error. Try to cast to a CcaError for more details
	//Might not work if connectivity error or some other unexpected error
	if err != nil {
		if errorResponse, ok := err.(api.CcaErrorResponse); ok {
			if errorResponse.StatusCode == api.NOT_FOUND {
				fmt.Println("Volume was not found")
			} else {
				//Can get more details from the CcaErrors
				fmt.Println(errorResponse.Errors)
			}
		} else {
			//handle unexpected error
		}
	}
}
```


#License

This project is licensed under the terms of the MIT license.
