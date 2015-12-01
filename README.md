# go-cloudca

Cloud.ca client for the Go programming language

# How to use

First of all create a new CcaClient.
```
	ccaClient := gocca.NewCcaClient("[your-api-key]")
```

Get the ServiceResources object for a specific environment and service. Here, we assume that it is a cloudca service.
```
	resources, _ := ccaClient.GetResources("[service-code]", "[environment-name]")
	ccaResources := resources.(cloudca.Resources)
```

Now with the cloudca ServiceResources object, we can execute operations on cloudca resources in the specified environment.

Retrieve the list of instances in the environment.
```
	instances, _ := ccaResources.Instances.List()
```

Get a specific volume in the environment.
```
	volume, _ := ccaResources.Volumes.Get("[some-volume-id]")
```

Create a new instance in the environment.
```
	createdInstance, _ := ccaResources.Instances.Create(cloudca.Instance{
			Name: "[new-instance-name]",
			TemplateId: "[some-template-id]",
			ComputeOfferingId:"[some-compute-offering-id]",
			NetworkId:"[some-network-id]",
		})
```

#Handling Errors

```
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
```


#License

This project is licensed under the terms of the MIT license.
