## S3 CLI

A CLI for uploading, fetching and deleting content to/from an S3-compatible
blobstore.

Continuous integration: <https://github.com/cloudfoundry/storage-cli/actions/workflows/s3-integration.yml>

Releases can be found in `https://s3.amazonaws.com/bosh-s3cli-artifacts`. The Linux binaries follow the regex `s3cli-(\d+\.\d+\.\d+)-linux-amd64` and the windows binaries `s3cli-(\d+\.\d+\.\d+)-windows-amd64`.

## Usage

Given a JSON config file (`config.json`)...

``` json
{
  "bucket_name":            "<string> (required)",

  "credentials_source":     "<string> [static|env_or_profile|none]",
  "access_key_id":          "<string> (required if credentials_source = 'static')",
  "secret_access_key":      "<string> (required if credentials_source = 'static')",

  "region":                 "<string> (optional - default: 'us-east-1')",
  "host":                   "<string> (optional)",
  "port":                   <int> (optional),

  "ssl_verify_peer":        <bool> (optional),
  "use_ssl":                <bool> (optional),
  "signature_version":      "<string> (optional)",
  "server_side_encryption": "<string> (optional)",
  "sse_kms_key_id":         "<string> (optional)",
  "multipart_upload":       <bool> (optional - default: true)
}
```

``` bash
# Usage
s3-cli --help

# Command: "put"
# Upload a blob to an S3-compatible blobstore.
s3-cli -c config.json put <path/to/file> <remote-blob>

# Command: "get"
# Fetch a blob from an S3-compatible blobstore.
# Destination file will be overwritten if exists.
s3-cli -c config.json get <remote-blob> <path/to/file>

# Command: "delete"
# Remove a blob from an S3-compatible blobstore.
s3-cli -c config.json delete <remote-blob>

# Command: "exists"
# Checks if blob exists in an S3-compatible blobstore.
s3-cli -c config.json exists <remote-blob>

# Command: "sign"
# Create a self-signed url for an object
s3-cli -c config.json sign <remote-blob> <get|put> <seconds-to-expiration>
```

## Contributing

Follow these steps to make a contribution to the project:

- Fork this repository
- Create a feature branch based upon the `main` branch (*pull requests must be made against this branch*)
  ``` bash
  git checkout -b feature-name origin/main
  ```
- Run tests to check your development environment setup
  ``` bash
  ginkgo --race --skip-package=integration --randomize-all --cover -v -r ./s3/...
  ```
- Make your changes (*be sure to add/update tests*)
- Run tests to check your changes
  ``` bash
  ginkgo --race --skip-package=integration --randomize-all --cover -v -r ./s3/...
  ```
- Push changes to your fork
  ``` bash
  git add .
  git commit -m "Commit message"
  git push origin feature-name
  ```
- Create a GitHub pull request, selecting `main` as the target branch

## Testing

### Unit Tests
**Note:** Run the following commands from the repository root directory.
  ``` bash
  go install github.com/onsi/ginkgo/v2/ginkgo

  ginkgo --skip-package=integration --randomize-all --cover -v -r ./s3/...
  ```

### Integration Tests

To run the integration tests, export the following variables into your environment:

```
export access_key_id=YOUR_AWS_ACCESS_KEY
export focus_regex="GENERAL AWS|AWS V2 REGION|AWS V4 REGION|AWS US-EAST-1"
export region_name=us-east-1
export s3_endpoint_host=https://s3.amazonaws.com
export secret_access_key=YOUR_SECRET_ACCESS_KEY
export stack_name=s3cli-iam
export bucket_name=s3cli-pipeline
```

Run `./.github/scripts/s3/setup-aws-infrastructure.sh` and `./.github/scripts/s3/teardown-infrastructure.sh` before and after the `./.github/scripts/s3/run-integration-*` in repo's root folder.
