package mgtt

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMgttAzurermRg() *schema.Resource {
	return &schema.Resource{
		Create: resourceMgttAzurermRgCreate,
		Read:   resourceMgttAzurermRgRead,
		Update: resourceMgttAzurermRgUpdate,
		Delete: resourceMgttAzurermRgDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceMgttAzurermRgCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	location := d.Get("location").(string)

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")

	azureResourceGroupHandler := NewAzureResourceGroupHandler(subscriptionID, accessToken)

	createRequestBody := map[string]interface{}{
		"location": location,
	}

	jsonString, err := ConvertMapToJSON(createRequestBody)
	if err != nil {
		return fmt.Errorf("Error converting map to JSON: %s", err)
	}

	err = azureResourceGroupHandler.CreateResourceGroup(name, jsonString)

	if err != nil {
		return err
	}

	id := uuid.New()
	d.SetId(id.String())

	if err := d.Set("name", name); err != nil {
		return err
	}
	if err := d.Set("location", location); err != nil {
		return err
	}
	return nil
}

func resourceMgttAzurermRgRead(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")
	azureResourceGroupHandler := NewAzureResourceGroupHandler(subscriptionID, accessToken)

	err := azureResourceGroupHandler.GetResourceGroup(name)

	if err != nil {
		return err
	}
	return nil
}

func resourceMgttAzurermRgUpdate(d *schema.ResourceData, m interface{}) error {
	oldName, oldNameExists := d.GetChange("name")
	oldLocation, oldLocationExists := d.GetChange("location")

	if oldNameExists.(bool) {
		fmt.Printf("Old name: %s\n", oldName.(string))
	} else {
		return fmt.Errorf("Error retrieving old name")
	}

	if oldLocationExists.(bool) {
		fmt.Printf("Old location: %s\n", oldLocation.(string))
	} else {
		return fmt.Errorf("Error retrieving old location")
	}

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")

	azureResourceGroupHandler := NewAzureResourceGroupHandler(subscriptionID, accessToken)

	err := azureResourceGroupHandler.DeleteResourceGroup(oldName.(string))
	if err != nil {
		return err
	}

	name := d.Get("name").(string)
	location := d.Get("location").(string)
	createRequestBody := map[string]interface{}{
		"location": location,
	}

	jsonString, err := ConvertMapToJSON(createRequestBody)
	if err != nil {
		return fmt.Errorf("Error converting map to JSON: %s", err)
	}

	err = azureResourceGroupHandler.CreateResourceGroup(name, jsonString)

	if err != nil {
		return err
	}

	if err := d.Set("name", name); err != nil {
		return err
	}
	if err := d.Set("location", location); err != nil {
		return err
	}

	return nil
}

func resourceMgttAzurermRgDelete(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")
	azureResourceGroupHandler := NewAzureResourceGroupHandler(subscriptionID, accessToken)

	err := azureResourceGroupHandler.DeleteResourceGroup(name)

	if err != nil {
		return err
	}
	return nil
}
