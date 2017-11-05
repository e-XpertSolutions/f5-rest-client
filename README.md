# f5-rest-client

[![GoDoc](https://godoc.org/github.com/e-XpertSolutions/f5-rest-client/f5?status.png)](http://godoc.org/github.com/e-XpertSolutions/f5-rest-client/f5)
[![License](https://img.shields.io/badge/license-BSD%203--Clause-yellow.svg?style=flat)](https://github.com/e-XpertSolutions/f5-rest-client/blob/master/LICENSE)

f5-rest-client implements a REST client to query the F5 Big IP API.


## Installation

```
go get -u github.com/e-XpertSolutions/f5-rest-client/f5
```


## Usage

```go
package main

import (
	"log"

	"github.com/e-XpertSolutions/f5-rest-client/f5"
	"github.com/e-XpertSolutions/f5-rest-client/f5/ltm"
)

func main() {
	// setup F5 BigIP client
	f5Client, err := f5.NewBasicClient("https://url-to-bigip", "user", "password")
	if err != nil {
		log.Fatal(err)
	}

	// setup client for the LTM API
	ltmClient := ltm.New(f5Client)

	// query the /ltm/virtual API
	vsConfigList, err := ltmClient.Virtual().ListAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(vsConfigList)
}
```


## Features

- [x] Add support for HTTP Basic Authentication
- [x] Add support for token based authentication
- [ ] Add support for authentication through external providers
- [x] Manage Virtual Server, pool, node, irules, monitors (/ltm)
- [x] Manage Cluster Management (/cm)
- [x] Manage interfaces, vlan, trunk, self ip, route, route domains (/net)
- [x] Manage system related stuffs (/sys)
- [ ] Manage firewall, WAF and DOS profiles (/security)
- [ ] Manage virtualization features (/vcmp)
- [ ] Manage access policies (/apm)
- [x] Manage DNS and global load balancing servers (/gtm)
- [ ] Add support for analytics read-only API (/analytics)
- [ ] Add support for results pagination
- [x] Add support for transaction

## Contributing

We appreciate any form of contribution (feature request, bug report,
pull request, ...). We have no special requirements for Pull Request,
just follow the standard [GitHub way](https://help.github.com/articles/using-pull-requests/).


## License

The sources are release under a BSD 3-Clause License. The full terms of that
license can be found in `LICENSE` file of this repository.
