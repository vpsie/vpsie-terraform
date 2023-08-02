package template

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Here we define da linked list of all the resources that we want to
// support in our provider. As an example, if you were to write an AWS provider
// which supported resources like ec2 instances, elastic balancers and things of that sort
// then this would be the place to declare them.
func ResourceIP() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createIPFunc,
		Read:          readIPFunc,
		Update:        updateIPFunc,
		Delete:        deleteIPFunc,
		Schema: map[string]*schema.Schema{ // List of supported configuration fields for your resource
			"vmIdentifier": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ipType": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

func createIPFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*VPSieClient)
	ip := IP{
		VmIdentifier: d.Get("vmidentifier").(string),
		IpType:       d.Get("iptype").(string),
	}

	err := client.CreateIP(&ip)
	if err != nil {
		return err
	}

	// todo set the ID
	ip.Identifier = "123"
	d.SetId(ip.Identifier)

	return nil
}

func readIPFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*VPSieClient)
	id := d.Id()
	err := client.ReadIP(id)
	if err != nil {
		return err
	}
	_ = IP{
		Identifier: id,
	}

	return nil
}

func updateIPFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteIPFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*VPSieClient)
	id := d.Id()
	ip := IP{
		Identifier: id,
	}

	err := client.DeleteIP(&ip)
	if err != nil {
		return err
	}

	return nil
}
