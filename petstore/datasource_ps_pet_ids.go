package petstore

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sdk "github.com/terraform-in-action/go-petstore"
)

func dataSourcePSPets() *schema.Resource {
	return &schema.Resource{
		Read: resourcePSPetRead,

		Schema: map[string]*schema.Schema{
			"names": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourcePSPetRead(d *schema.ResourceData, meta interface{}) error {
	names := make(map[string]bool)
	for _, name := range d.Get("names").([]interface{}) {
		id += name.(string)
		names[name.(string)] = true
	}

	conn := meta.(*sdk.Client)
	pets, err := conn.Pets.List()
	if err != nil {
		return err
	}

	var ids []string
	for _, pet := range pets {
		if names["*"] || names[pet.Name] {
			ids = append(ids, pet.ID)
		}
	}
	d.Set("ids", ids)
	return nil
}
