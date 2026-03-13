<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset=".github/cloudbats-lockup-horizontal-dark-960.png">
    <img src=".github/cloudbats-lockup-horizontal-light-960.png" alt="CloudBats" width="400">
  </picture>
</p>

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

| Version | Status |
|---------|--------|
| Network Application v10.1.85 (UX7) | Active development and testing |
| UniFi Controller v8.4.59 | Previously supported |
| UniFi Controller up to v7.4.162 | Supported in original [paultyng](https://github.com/paultyng/terraform-provider-unifi) release |

## Development Status

This provider is under active development targeting the **UniFi Express 7 (UX7)** running **Network Application v10.1.85**.

### Tested and Working (v10.1.85)

| Resource / Data Source | Status | Notes |
|------------------------|--------|-------|
| `unifi_network` | **Working** | Create, update, import. `internet_access_enabled` fix applied. |
| `unifi_firewall_group` | **Working** | Port groups for firewall rules. |
| `unifi_dynamic_dns` | **Working** | Tested with DuckDNS. |
| `data.unifi_user_group` | **Working** | Default user group lookup. |
| `data.unifi_ap_group` | **Working** | Default AP group lookup. |
| Multi-site (aliased providers) | **Working** | Separate provider blocks for multiple UX7 controllers. |

### Known Limitations (v10.1.85)

| Resource / Feature | Status | Notes |
|--------------------|--------|-------|
| `unifi_firewall_rule` | **Not working** | UX7 uses Zone-Based Firewall engine; returns `api.err.FirewallRuleIndexOutOfRange`. |
| WireGuard VPN | **Not supported** | No Terraform resource; must be configured via UI. |
| Traffic Routes / Policy-Based Routing | **Not supported** | No Terraform resource; must be configured via UI. |
| Zone-Based Firewall | **Not supported** | New API in Network 10.x+, not yet mapped. |

### Not Yet Tested on v10.1.85

| Resource | Notes |
|----------|-------|
| `unifi_port_profile` | Previously working on v8.4.59, untested on v10.1.85. |
| `unifi_wlan` | WLAN/SSID management, untested. |
| `unifi_user` | Static client/user management, untested. |
| `unifi_device` | Device adoption/config, untested. |
| `unifi_port_forward` | Port forwarding rules, untested. |
| `unifi_radius_profile` | RADIUS authentication, untested. |
| `unifi_account` | RADIUS account management, untested. |
| `unifi_setting_mgmt` | Management settings, untested. |
| `unifi_setting_usg` | USG/gateway settings, untested. |

### Key Fixes in This Fork

- **`internet_access_enabled` schema re-enabled**: Was commented out as "deprecated" but the Go struct still sent `false` to the API (no `omitempty`), silently disabling internet for Terraform-created networks.
- **JSON unmarshaling fix (go-unifi)**: UniFi API `.data` field returns either an object or array; fixed to handle both with `json.RawMessage`.
- **Local admin auth**: Bypasses SSO 2FA for API access using a local-only admin account.

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
