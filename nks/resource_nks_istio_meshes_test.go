package nks

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccResourceIstioMesh_Basic(t *testing.T) {
	_, exists := os.LookupEnv("TF_ACC_ISTIO_MESH_LOCK")
	if !exists {
		t.Skip("`TF_ACC_ISTIO_MESH_LOCK` isn't specified - skipping since test will increase test time significantly")
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testResourceNKSIstioMesh_Basic,
			},
		},
	})
}

const testResourceNKSIstioMesh_Basic = `
data "nks_organization" "org" {

}

data "nks_keyset" "keyset_default" {
	org_id   = "${data.nks_organization.org.id}"
	name     = "new Azure creds"
	category = "provider"
	entity   = "azure"
}

data "nks_keyset" "ssh" {
	org_id   = "${data.nks_organization.org.id}"
	category = "user_ssh"
	name = "Default SPC SSH Keypair"
}

data "nks_instance_specs" "master-specs" {
	provider_code = "azure"
	node_size     = "standard_f1"
}

data "nks_instance_specs" "worker-specs" {
	provider_code = "azure"
	node_size     = "${data.nks_instance_specs.master-specs.node_size}"
}

resource "nks_cluster" "terraform-cluster-a" {
	org_id                  = "${data.nks_organization.org.id}"
	cluster_name            = "TF IstioMeshAcceptance ClusterA 1"
	provider_code           = "azure"
	provider_keyset         = "${data.nks_keyset.keyset_default.id}"
	region                  = "eastus"
	k8s_version             = "v1.13.2"
	startup_master_size     = "${data.nks_instance_specs.master-specs.node_size}"
	startup_worker_count    = 2
	startup_worker_size     = "${data.nks_instance_specs.worker-specs.node_size}"
	provider_network_cidr   = "10.0.0.0/16"
	provider_subnet_cidr    = "10.0.0.0/24"
	rbac_enabled            = true
	dashboard_enabled       = true
	etcd_type               = "classic"
	platform                = "coreos"
	channel                 = "stable"
	timeout                 = 1800
	ssh_keyset              = "${data.nks_keyset.ssh.id}"
}

resource "nks_cluster" "terraform-cluster-b" {
	org_id                  = "${data.nks_organization.org.id}"
	cluster_name            = "TF IstioMeshAcceptance ClusterA 2"
	provider_code           = "azure"
	provider_keyset         = "${data.nks_keyset.keyset_default.id}"
	region                  = "eastus"
	k8s_version             = "v1.13.2"
	startup_master_size     = "${data.nks_instance_specs.master-specs.node_size}"
	startup_worker_count    = 2
	startup_worker_size     = "${data.nks_instance_specs.worker-specs.node_size}"
	provider_network_cidr   = "10.0.0.0/16"
	provider_subnet_cidr    = "10.0.0.0/24"
	rbac_enabled            = true
	dashboard_enabled       = true
	etcd_type               = "classic"
	platform                = "coreos"
	channel                 = "stable"
	timeout                 = 1800
	ssh_keyset              = "${data.nks_keyset.ssh.id}"
  }

resource "nks_solution" "istio-a" {
	org_id     = "${data.nks_organization.org.id}"
	cluster_id = "${nks_cluster.terraform-cluster-a.id}"
	solution   = "istio"
}

resource "nks_solution" "istio-b" {
	org_id     = "${data.nks_organization.org.id}"
	cluster_id = "${nks_cluster.terraform-cluster-b.id}"
	solution   = "istio"
}

resource "nks_istio_mesh" "istio-mesh-tf" {
	name        = "tf-istio-mesh-1"
	mesh_type   = "cross_cluster"
	org_id      = "${data.nks_organization.org.id}"
	workspace   = "105"
	members     = [
		{
			cluster	= "${nks_cluster.terraform-cluster-a.id}"
			role	= "host"
		},
		{
			cluster	= "${nks_cluster.terraform-cluster-b.id}"
			role	= "guest"			
		}
	]
}
`
