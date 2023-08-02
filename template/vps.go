package template

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Here we define da linked list of all the resources that we want to
// support in our provider. As an example, if you were to write an AWS provider
// which supported resources like ec2 instances, elastic balancers and things of that sort
// then this would be the place to declare them.
func ResourceVPS() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVPSFunc,
		Read:          readVPSFunc,
		Update:        updateVPSFunc,
		Delete:        deleteVPSFunc,
		Schema: map[string]*schema.Schema{ // List of supported configuration fields for your resource
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"resourceIdentifier": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"osIdentifier": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"dcIdentifier": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"notes": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"backupEnabled": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"addPublicIpV4": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"addPublicIpV6": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"addPrivateIp": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sshKeyIdentifier": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": &schema.Schema{
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

func createVPSFunc(d *schema.ResourceData, meta interface{}) error {
	client, err := CreateVPSieClientFromResourceData(d)
	if err != nil {
		return err
	}
	vps := VPS{
		Hostname: d.Get("hostname").(string),
	}

	err = client.CreateVPS(&vps)
	if err != nil {
		return err
	}

	d.SetId(vps.Id())

	return nil
}

func readVPSFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func updateVPSFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteVPSFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}
