package template

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Here we define da linked list of all the resources that we want to
// support in our provider. As an example, if you were to write an AWS provider
// which supported resources like ec2 instances, elastic balancers and things of that sort
// then this would be the place to declare them.
func ResourceStorage() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createStorageFunc,
		Read:          readStorageFunc,
		Update:        updateStorageFunc,
		Delete:        deleteStorageFunc,
		Schema: map[string]*schema.Schema{ // List of supported configuration fields for your resource
			"storages": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}

}

// The methods defined below will get called for each resource that needs to
// get created (createFunc), read (readFunc), updated (updateFunc) and deleted (deleteFunc).
// For example, if 10 resources need to be created then `createFunc`
// will get called 10 times every time with the information for the proper
// resource that is being mapped.
//
// If at some point any of these functions returns an error, Terraform will
// imply that something went wrong with the modification of the resource and it
// will prevent the execution of further calls that depend on that resource
// that failed to be created/updated/deleted.

func createStorageFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*VPSieClient)
	storages := d.Get("storages").(map[string]string)
	storage := Storage{
		Items: storages,
	}

	err := client.CreateStorage(&storage)
	if err != nil {
		return err
	}

	// set the ID
	storage.Identifier = "123"
	d.SetId(storage.Identifier)

	return nil
}

func readStorageFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*VPSieClient)
	id := d.Id()
	err := client.ReadStorage(id)
	if err != nil {
		return err
	}
	_ = Storage{
		Identifier: id,
	}

	return nil

}

func updateStorageFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteStorageFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*VPSieClient)
	id := d.Id()
	storage := Storage{
		Identifier: id,
	}

	err := client.DeleteStorage(&storage)
	if err != nil {
		return err
	}

	return nil

}
