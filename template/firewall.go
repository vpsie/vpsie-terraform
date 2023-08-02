package template

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Here we define da linked list of all the resources that we want to
// support in our provider. As an example, if you were to write an AWS provider
// which supported resources like ec2 instances, elastic balancers and things of that sort
// then this would be the place to declare them.
func ResourceFirewall() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createFirewallFunc,
		Read:          readFirewallFunc,
		Update:        updateFirewallFunc,
		Delete:        deleteFirewallFunc,
		Schema: map[string]*schema.Schema{ // List of supported configuration fields for your resource
			"groupName": &schema.Schema{
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

func createFirewallFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*VPSieClient)
	firewall := Firewall{
		Hostname: d.Get("hostname").(string),
	}

	err := client.CreateFirewall(&firewall)
	if err != nil {
		return err
	}

	d.SetId(firewall.Id())

	return nil
}

func readFirewallFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func updateFirewallFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteFirewallFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}
