# GCE Cluster example

This example will show you how to create a cluster on GCE with Netapp Kubernetes Services.

This example does the following:
- Finds organization
- Finds GCE keyset
- Finds SSH keyset
- Creates a cluster

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