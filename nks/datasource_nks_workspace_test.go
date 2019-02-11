package nks

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceWorkspace_Basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceNKSWorkspace_Basic,
			},
		},
	})
}

const testAccDataSourceNKSWorkspace_Basic = `
resource "nks_workspace" "my-workspace" {
    org_id          = "159"
	default         = false
	name            = "My Test Workspace"
}
`
