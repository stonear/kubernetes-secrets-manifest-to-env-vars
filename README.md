# kubernetes-secrets-manifest-to-env-vars
 
Convert Kubernetes secrets manifests to environment variables.

## Usage

```bash
go run . -help
```

## Example

```bash
go run . -input example.yaml -output example.env.sh
```

## Options

- `-input` The path to the YAML file containing the Kubernetes secrets manifests.
- `-output` The path to the output file.

## Contributing

Contributions are welcome. Please open a pull request with your changes.
