package nks

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceKeyset_Basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceNKSKeyset_Basic,
			},
		},
	})
}

const testAccDataSourceNKSKeyset_Basic = `
data "nks_organization" "default" {
    name = "My Organization"
}

resource "nks_keyset" "keyset-default" {
    org_id = "${data.nks_organization.default.id}"
    name = "My default SSH"
    category = "user_ssh"
    keys = [
        {
            key_type = "pub"
            key = "dummyPubKey=="
        }
    ]
}
`
