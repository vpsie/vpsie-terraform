package template

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

type VPSieClient struct {
	ClientID  string
	SecretKey string
}

type VPS struct {
	ResourceIdentifier string   `json:""`
	OsIdentifier       string   `json:""`
	DcIdentifier       string   `json:""`
	Hostname           string   `json:""`
	Notes              string   `json:""`
	BackupEnabled      bool     `json:""`
	AddPublicIpV4      bool     `json:""`
	AddPublicIpV6      bool     `json:""`
	AddPrivateIp       bool     `json:""`
	SshKeyIdentifier   string   `json:""`
	Tags               []string `json:""`
}

type Firewall struct {
	ResourceIdentifier string   `json:""`
	OsIdentifier       string   `json:""`
	DcIdentifier       string   `json:""`
	Hostname           string   `json:""`
	Notes              string   `json:""`
	BackupEnabled      bool     `json:""`
	AddPublicIpV4      bool     `json:""`
	AddPublicIpV6      bool     `json:""`
	AddPrivateIp       bool     `json:""`
	SshKeyIdentifier   string   `json:""`
	Tags               []string `json:""`
}

type FirewallRecord struct {
	FirewallId string `json:""`
}

type Snapshot struct {
	Identifier   string `json:""`
	VmIdentifier string `json:""`
	Name         string `json:""`
}

type Image struct {
	DcIdentifier string `json:""`
	ImageName    string `json:""`
	ImageUrl     string `json:""`
}

type Backup struct {
	Identifier   string `json:""`
	VmIdentifier string `json:""`
	Name         string `json:""`
	Notes        string `json:""`
}

type IP struct {
	Identifier   string `json:""`
	VmIdentifier string `json:""`
	IpType       string `json:""`
}

type FIP struct {
	Identifier   string `json:""`
	VmIdentifier string `json:""`
	DcIdentifier string `json:""`
	IpType       string `json:""`
}

type SSHKey struct {
	Identifier string `json:""`
	PrivateKey string `json:""`
	Name       string `json:""`
}
type Storage struct {
	Identifier string            `json:""`
	Items      map[string]string `json:""`
}

type LoadBalancer struct {
	Identifier string `json:""`
}

func (m *VPS) Id() string {
	return "id-" + m.Hostname + "!"
}

func (m *Firewall) Id() string {
	return "id-" + m.Hostname + "!"
}

func (m *FirewallRecord) Id() string {
	return "id-firewall"
}

func (m *Image) Id() string {
	return "id-image"
}

func CreateVPSieClientFromResourceData(d *schema.ResourceData) (*VPSieClient, error) {
	client := VPSieClient{
		ClientID:  d.Get("client_id").(string),
		SecretKey: d.Get("secret_api_key").(string),
	}
	return &client, nil
}

func (c *VPSieClient) getAuthToken() (string, error) {
	params := struct {
		ClientId     string `json:"clientId"`
		ClientSecret string `json:"clientSecret"`
	}{
		ClientId:     c.ClientID,
		ClientSecret: c.SecretKey,
	}
	body, err := json.Marshal(params)

	if err != nil {
		return "", err
	}
	fmt.Println("response: " + string(body))

	_, err = SendHttpRequest("/auth/from/api", body)
	return "", nil
}

func (c *VPSieClient) CreateVPS(m *VPS) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) CreateFirewall(m *Firewall) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) CreateFirewallRecord(m *FirewallRecord) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) CreateSnapshot(m *Snapshot) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) ReadSnapshot(id string) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) DeleteSnapshot(m *Snapshot) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) CreateImage(m *Image) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) CreateBackup(m *Backup) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) ReadBackup(id string) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) DeleteBackup(m *Backup) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) CreateSSHKey(m *SSHKey) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) DeleteSSHKey(m *SSHKey) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) ReadSSHKey(id string) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) CreateIP(m *IP) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) DeleteIP(m *IP) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) ReadIP(id string) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) CreateFIP(m *FIP) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) ReadFIP(id string) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) DeleteFIP(m *FIP) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) CreateStorage(m *Storage) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) DeleteStorage(m *Storage) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) ReadStorage(id string) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) CreateLoadBalancer(m *LoadBalancer) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) CreateLoadBalancerRules(m *LoadBalancer, rules map[string]string) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) DeleteLoadBalancer(m *LoadBalancer) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}

func (c *VPSieClient) ReadLoadBalancer(id string) error {
	_, err := c.getAuthToken()
	if err != nil {
		return err
	}

	return nil
}
