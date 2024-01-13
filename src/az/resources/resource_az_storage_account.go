package terraform_provider_blob_services

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAzStorageAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceAzStorageAccountCreate,
		Read:   resourceAzStorageAccountRead,
		Update: resourceAzStorageAccountUpdate,
		Delete: resourceAzStorageAccountDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAzStorageAccountCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAzStorageAccountRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAzStorageAccountUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAzStorageAccountDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
