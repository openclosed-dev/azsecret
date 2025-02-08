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

### Building with installed Go

On Ubuntu 24.04 LTS, the following command installs the tools required for development.

```
sudo apt install golang-go build-essential
```

Invoke the following command in the repository root.
```
make
```

### Building with Docker

Build a new Docker image for Golang development using the Dockerfile in this repository root.
```
docker image build -t go-builder:latest .
```

And run a Docker container from the image.
```
docker container run --rm -v '.:/repo' go-builder:latest
```
