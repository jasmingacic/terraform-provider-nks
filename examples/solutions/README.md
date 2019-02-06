# AKS Cluster example

This example will show you how to create a cluster on AKS with Netapp Kubernetes Services.

This example does the following:
- Finds organization
- Finds AKS keyset
- Finds SSH keyset
- Creates a cluster

## Run the example

From inside of this directory:

```bash
export NKS_API_TOKEN=<this is a secret>
export NKS_API_URL=<this is a secret>
terraform init
terraform plan
terraform apply
```

## Remove the example

```bash
terraform destroy