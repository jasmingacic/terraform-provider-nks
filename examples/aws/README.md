# AWS Cluster example

This example will show you how to create a cluster on AWS with Netapp Kubernetes Services.

This example does the following:
- Finds organization
- Finds AWS keyset
- Finds SSH keyset
- Creates a cluster
- Add an additional master node
- Add an additional node pool

[Keyset examples](/examples/keysets) shows how to add a key to NKS.

## Run the example

From inside of this directory:

```bash
export NKS_API_TOKEN=<this is a secret>
export NKS_API_URL=https://api.stackpoint.io/
terraform init
terraform plan
terraform apply
```

## Remove the example

```bash
terraform destroy