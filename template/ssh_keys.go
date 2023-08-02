package template

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Here we define da linked list of all the resources that we want to
// support in our provider. As an example, if you were to write an AWS provider
// which supported resources like ec2 instances, elastic balancers and things of that sort
// then this would be the place to declare them.
func ResourceSSHKey() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSSHKeyFunc,
		Read:          readSSHKeyFunc,
		Update:        updateSSHKeyFunc,
		Delete:        deleteSSHKeyFunc,
		Schema: map[string]*schema.Schema{ // List of supported configuration fields for your resource
			"privatekey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
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

func createSSHKeyFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*VPSieClient)
	sshKey := SSHKey{
		PrivateKey: d.Get("privatekey").(string),
		Name:       d.Get("name").(string),
	}

	err := client.CreateSSHKey(&sshKey)
	if err != nil {
		return err
	}

	id := sshKey.Identifier
	d.SetId(id)

	return nil
}

func readSSHKeyFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*VPSieClient)
	id := d.Id()
	err := client.ReadSSHKey(id)
	if err != nil {
		return err
	}
	_ = SSHKey{
		Name: id,
	}

	return nil
}

func updateSSHKeyFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteSSHKeyFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*VPSieClient)
	id := d.Id()
	key := SSHKey{
		Identifier: id,
	}

	err := client.DeleteSSHKey(&key)
	if err != nil {
		return err
	}

	return nil

}
