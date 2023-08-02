# Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

# Usage

```
# For example, restrict template version in 0.1.x
provider "vpsie" {
  version = "~> 0.1"
}
```

# Building The Provider

Clone repository to: `$GOPATH/src/github.com/vpsie/terraform-provider`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone https://github.com/vpsie/terraform-provider
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/vpsie/terraform-provider
$ make build
```

# Using the provider

You may want to use the provider to create a VPS or in more advanced scenarios you may want to create a full stack including a firewall and backup routine. Below are some basic examples to help you get started:

## Creating a VPS instance

The following examples creates a basic VPS

```
vps {
  resourceIdentifier = "9a0e4e84-9f22-11e3-8af5-005056aa8af7"
  osIdentifier = "75401d7d-d9d3-11e3-b135-005056aa8af7"
  dcIdentifier = "ab70b3b6-9f22-11e3-8af5-005056aa8af7"
  hostname = "test-hostIP"
  notes = "notes"
  backupEnabled = false
  addPublicIpV4 = false
  addPublicIpV6 = false
  addPrivateIp" = false
  sshKeyIdentifier = "4747227d-af42-11e6-a5af-005056aadd24"
  tags = [
    "tag1"
  ]
}
```

# Development

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-template
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

# Contact

If you need more information or have any questions please reach out to VPSie developers