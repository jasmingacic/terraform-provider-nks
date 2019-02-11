# Organization
organization_name = "Wandering Band"

# Cluster
a_cluster_name = "TF istio meshes Cluster A"
b_cluster_name = "TF istio meshes Cluster B"

# Keyset
provider_keyset_name = "new Azure creds"
ssh_keyset_name = "Default SPC SSH Keypair"

# Workspace
workspace_id = 105

# Provider
provider_code = "azure"
provider_k8s_version = "v1.13.2"
provider_platform = "coreos"
provider_region = "eastus"
provider_resource_group = "__new__"
provider_network_id = "__new__"
provider_network_cidr = "10.0.0.0/16"
provider_subnet_id = "__new__"
provider_subnet_cidr = "10.0.1.0/24"
provider_master_size = "standard_f1"
provider_worker_size = "standard_f1"
provider_channel = "stable"
provider_etcd_type = "classic"

# Istio mesh
istio_mesh_name = "tf-istio-mesh"
istio_mesh_type = "cross_cluster"