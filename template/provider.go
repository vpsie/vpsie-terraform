package template

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: providerSchema(),
		DataSourcesMap: map[string]*schema.Resource{
			"template_file":             dataSourceFile(),
			"template_cloudinit_config": dataSourceCloudinitConfig(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"vps":      ResourceVPS(),
			"sshkey":   ResourceSSHKey(),
			"ip":       ResourceIP(),
			"firewall": ResourceFirewall(),
			"image":    ResourceImage(),
			"snapshot": ResourceSnapshot(),
		},
	}
}

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"secret_api_key": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "VPSie secret API key",
		},
		"client_id": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "VPSie client ID",
		},
	}
}
