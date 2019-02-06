# Keysets example

The following examples will show you how to add a keyset to Netapp Kubernetes Services. Keysets will allow NKS to do all necessary operations on one or more cloud provides so your cluster could be operational.  

## Adding cloud provider keyset

```bash
data "nks_organization" "default" {
  name = "My Organization"
}

resource "nks_keyset" "default_aks" {
     org_id="${data.nks_organization.default.id}"
     name="My default AKS"
     category = "provider"
     entity = "aks"
     keys = [
         {
             key_type = "pub"
             key = "MY_PUB_KEY"
         },
         {
             key_type = "pvt"
             key = "MY_PRIVATE_KEY"
         },
     ]
}
```

## Adding SSH keyset

```bash
data "nks_organization" "default" {
  name = "My Organization"
}

resource "nks_keyset" "default_ssh" {
     org_id="${data.nks_organization.default.id}"
     name="My default SSH"
     category = "user_ssh"
     keys = [
         {
             key_type = "pub"
             key = "MY_PUB_KEY"
         }
     ]
}
```