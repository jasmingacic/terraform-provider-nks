package nks

import (
	"fmt"
	"log"
	"strconv"

	"github.com/StackPointCloud/stackpoint-sdk-go/stackpointio"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceNKSKeyset() *schema.Resource {
	return &schema.Resource{
		Create: resourceNKSKeysetCreate,
		Read:   resourceNKSKeysetRead,
		Delete: resourceNKSKeysetDelete,
		Schema: map[string]*schema.Schema{
			"org_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"category": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					if v.(string) != "provider" && v.(string) != "user_ssh" && v.(string) != "solution" {
						errors = append(errors, fmt.Errorf("Category can be one of following 'provider', 'user_ssh' or 'solution'"))
					}
					return
				},
			},
			"entity": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"keys": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"key": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceNKSKeysetCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	orgID := d.Get("org_id").(int)
	name := d.Get("name").(string)
	category := d.Get("category").(string)
	req := stackpointio.Keyset{
		Org:        orgID,
		Name:       name,
		Category:   category,
		Workspaces: []int{},
	}

	rawKeys := d.Get("keys").([]interface{})
	req.Keys = make([]stackpointio.Key, len(rawKeys))

	if temp, ok := d.GetOk("entity"); ok {
		if category == "user_ssh" {
			return fmt.Errorf("when 'category' is set to '%s', 'entity' cannot be set", category)
		}
		req.Entity = temp.(string)
	}

	for i, v := range rawKeys {
		value := v.(map[string]interface{})
		req.Keys[i] = stackpointio.Key{
			Type:  value["key_type"].(string),
			Value: value["key"].(string),
		}
	}
	log.Printf("[DEBUG] req.Keys", req.Keys)

	keyset, err := config.Client.CreateKeyset(orgID, req)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(keyset.ID))

	return resourceNKSKeysetRead(d, meta)
}

func resourceNKSKeysetRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	orgID := d.Get("org_id").(int)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	keyset, err := config.Client.GetKeyset(orgID, id)
	if err != nil {
		return err
	}

	d.Set("name", keyset.Name)
	d.Set("category", keyset.Category)
	d.Set("entity", keyset.Entity)

	return nil
}

func resourceNKSKeysetDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	orgID := d.Get("org_id").(int)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	return config.Client.DeleteKeyset(orgID, id)
}
