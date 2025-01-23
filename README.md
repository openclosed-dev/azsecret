# azsecret

_azsecret_ is a command-line program for retrieving a secret stored in Azure Key Vault
from Azure virtual machines with Managed Identity assigned.

## Usage

```
Retrieves a secret value stored in Azure Key Vault

Usage:
  azsecret [secret name] [flags]

Flags:
  -h, --help               help for azsecret
  -i, --identity string    Client ID of the Azure Managed Identity.
                           Defaults to the value of environment variable AZ_MANAGED_IDENTITY.
  -k, --key-vault string   Name of the Azure Key Vault.
                           Defaults to the value of environment variable AZ_KEY_VAULT.
```

The following environment variables control the behavior of the command.

| Environment variables | Description |
| --- | --- |
| `AZ_MANAGED_IDENTITY` | Provides the default client ID of the Managed Identity |
| `AZ_KEY_VAULT` | Provides the default name of the Key Vault |

## Building from source

### Prerequisites

On Ubuntu 24.04 LTS, the following command installs the tools required for development.

```
sudo apt install build-essential golang-go
```

### Building all

Invoke the following command in the repository root.
```
make
```
