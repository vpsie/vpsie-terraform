package template

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Here we define da linked list of all the resources that we want to
// support in our provider. As an example, if you were to write an AWS provider
// which supported resources like ec2 instances, elastic balancers and things of that sort
// then this would be the place to declare them.
func ResourceBackup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createBackupFunc,
		Read:          readBackupFunc,
		Update:        updateBackupFunc,
		Delete:        deleteBackupFunc,
		Schema: map[string]*schema.Schema{ // List of supported configuration fields for your resource
			"dcIdentifier": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"backupName": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"backupUrl": &schema.Schema{
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

func createBackupFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*VPSieClient)
	backup := Backup{
		VmIdentifier: d.Get("vmidentifier").(string),
		Name:         d.Get("name").(string),
		Notes:        d.Get("notes").(string),
	}

	err := client.CreateBackup(&backup)
	if err != nil {
		return err
	}

	backup.Identifier = "123"
	d.SetId(backup.Identifier)

	return nil
}

func readBackupFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func updateBackupFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteBackupFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}
