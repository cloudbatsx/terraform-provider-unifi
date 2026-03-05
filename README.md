# Unifi Terraform Provider (terraform-provider-unifi)

A Terraform provider for managing Ubiquiti UniFi Network infrastructure, maintained by [cloudbats](https://cloudbats.com).

## About This Project

This provider is now maintained by cloudbats under the [`cloudbatsx`](https://github.com/cloudbatsx) GitHub organization. It wraps the [cloudbats Go SDK for UniFi](https://github.com/cloudbatsx/go-unifi) to provide declarative infrastructure management for UniFi Network controllers via Terraform.

## Lineage

This project has been through several maintainers:

1. **[paultyng/terraform-provider-unifi](https://github.com/paultyng/terraform-provider-unifi)** — Original creator. Supported UniFi Network up to v7.4.162. Development went stale.
2. **[sayedh/terraform-provider-unifi](https://github.com/sayedh/terraform-provider-unifi)** — Forked in August 2024 by Sayed Haque to update support for newer UniFi versions (v8.4.59). Published on the [Terraform Registry](https://registry.terraform.io/providers/sayedh/unifi/latest) with 4,100+ downloads.
3. **[cloudbatsx/terraform-provider-unifi](https://github.com/cloudbatsx/terraform-provider-unifi)** (this repo) — Migrated under cloudbats in March 2026 to provide long-term, organization-backed maintenance and continue development toward the latest UniFi versions.

## Documentation

Provider documentation will be available on the [Terraform Registry](https://registry.terraform.io/providers/cloudbatsx/unifi/latest/docs) once the first release is published.

## Supported UniFi Controller Versions

As of the latest release, this provider supports **UniFi Controller v8.4.59**. Earlier versions, up to **v7.4.162**, were supported in the original [paultyng](https://github.com/paultyng/terraform-provider-unifi) release.

## Development Status

This provider is under active development. The following resources are fully tested and confirmed working:

- **Resource: Network**
- **Resource: Port Profile**

Other resources and data sources may not yet be fully functional on the latest UniFi Network version. Testing is ongoing.

## Using the Provider

### Terraform Configuration Example

```hcl
terraform {
  required_providers {
    unifi = {
      source  = "cloudbatsx/unifi"
      version = ">= 1.0.0"
    }
  }
}

provider "unifi" {
  controller = "https://<unifi-controller-url>"
  username   = "admin"
  password   = "password"
}

resource "unifi_network" "example" {
  name = "Example Network"
  # Add relevant configuration options here
}
```

**Note**: When using this provider, ensure you're connected via a hard-wired connection to the UniFi Controller rather than WiFi, as configuring your network over a connection that could disconnect is risky.

### Migrating from sayedh/unifi

Update your `required_providers` block and run `terraform init -upgrade`:

```hcl
terraform {
  required_providers {
    unifi = {
      source  = "cloudbatsx/unifi"
      version = ">= 1.0.0"
    }
  }
}
```

If your state references the old provider, update it with:

```bash
terraform state replace-provider sayedh/unifi cloudbatsx/unifi
```

## Contributing

Contributions are welcome. Please submit issues and pull requests to the [GitHub repository](https://github.com/cloudbatsx/terraform-provider-unifi).

## Links

- [cloudbats.com](https://cloudbats.com) — consultancy home
- [cloudbatsx.com](https://www.cloudbatsx.com) — technical blog
- [cloudbats.ai](https://cloudbats.ai) — AI services
- [cloudbatsx/go-unifi](https://github.com/cloudbatsx/go-unifi) — underlying Go SDK
