# Account API
> Account Service is responsible for entire account section of job portal microservice

*   Create an account
*   Update an account
*   Delete an account
*   Suspend an account

## Setup

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
