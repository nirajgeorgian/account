# Account API
> Account Service is responsible for entire account section of job portal microservice

*   Create an account
*   Update an account
*   Delete an account
*   Suspend an account

## Setup
> Setup to start working on this project

### Install GoLang
[version as of writing: go version go1.12.4 darwin/amd64](https://golang.org/)

### setup `$GOPATH`
```bash
# In your bash profile
export GOPATH="/Users/nirajgeorgian/Documents/go"
export PATH=$PATH:$GOPATH/bin
```

> ### IMPORTANT! Make sure this repository is located at
```bash
$GOPATH/src/github.com/nirajgeorgian/account
```

### Install protobuf
Mac: `make setup-protobuf-mac`
Linux: `make setup-protobuf-linux`
>   See: [Error](http://google.github.io/proto-lens/installing-protoc.html) if there are any failures

### Setup Go environment

#### Install go dep tool (https://github.com/golang/dep)
```bash
make setup-dep
```

Install go dependencies*

```bash
make setup-go
```
> these need to be managed outside of the vendor/ directory because they are used in proto code generation

## Development
> run the api's locally

### Generate protos
> After updating protobuf files, you need to regenerate dependent code
```bash
make protos
```

### Build Services
```bash
make build
```


## running account service
> account service is build as an command line application.
> After running `make build` run `./bin/server --help` to view available commands
```bash
account service is responsible for CRUD with account entity

Usage:
  account [flags]
  account [command]

Available Commands:
  createAccount create an account with gRPC server on:3000
  help          Help about any command
  serve         serves the gRPC server
  version       Print the account service version

Flags:
  -c, --config string   config file (default is config.yaml)
  -h, --help            help for account
  -v, --verbose         verbose output (default is false)
  -r, --viper           Use Viper for configuration (default is true) (default true)

Use "account [command] --help" for more information about a command.
```
> use is like this `./bin/server [command] [flags]`
