# apicurio-cli
Command line tool for Apicurio projects.

## Installation

1. Please navigate to release and select binary file that matches your operating system
2. Install binary in the path

## Usage

```
[wtrocki@graphapi]$ apicr service-registry
Manage and interact with your Service Registry instances directly from the command line.

Create new Service Registry instances and interact with them by adding schema and API artifacts and downloading them to your computer.

Commands are divided into the following categories:

* Instance management commands: create, list, and so on
* Commands executed on selected instance: artifacts
* "use" command that selects the current instance

Usage:
  apicr service-registry [command]


Available Commands:
  artifact    Manage Service Registry artifacts
  create      Create a Service Registry instance
  delete      Delete a Service Registry instance
  describe    Describe a Service Registry instance
  list        List Service Registry instances
  role        Service Registry role management
  rule        Manage artifact rules in a Service Registry instance
  use         Use a Service Registry instance

Global Flags:
  -h, --help   Prints help information

Use "apicr service-registry [command] --help" for more information about a command.
```

## Using apicurio-cli with Operate First Apicurio Registries

To use apicurio-cli with Operate First Apicurio registries, use this login command:

```bash
apicr login --api-gateway https://fleet-manager-mt-apicurio-apicurio-registry.apps.smaug.na.operate-first.cloud --auth-url https://auth.apicur.io/auth/realms/operate-first-apicurio --client-id apicurio-cli
```

## Using apicurio-cli with Self Hosted Apicurio Registries

To use apicurio-cli with Self Hosted Apicurio registries, use this login command:

```bash
apicr login --api-gateway <fleet-manager-url> --auth-url <keycloak-url> --client-id <keycloak-relevant-client>
```

And bypass the checks while performing operations on the registries:

```bash
apicr service-registry create --bypass-checks
```

## Contributing

Check out the [Contributing Guide](./CONTRIBUTING.md) to learn more about the repository and how you can contribute.
