# f5-rest-client

[![GoDoc](https://godoc.org/github.com/e-XpertSolutions/f5-rest-client/f5?status.png)](http://godoc.org/github.com/e-XpertSolutions/f5-rest-client/f5)

f5-rest-client implements a REST client to query the F5 Big IP API.


## Installation

```
go get -u github.com/e-XpertSolutions/f5-rest-client/f5
```


## Usage


## Features

- [x] Add support for HTTP Basic Authentication
- [x] Add support for token based authentication
- [ ] Add support for authentication through external providers
- [x] Manage Virtual Server, pool, node, irules, monitors (/ltm)
- [x] Manage Cluster Management (/cm)
- [x] Manage interfaces, vlan, trunk, self ip, route, route domains (/net)
- [x] Manage system related stuffs (/sys)
- [ ] Manage virtualization features (/vcmp)
- [ ] Manage access policies (/apm)
- [ ] Manage DNS and global load balancing servers (/gtm)

## Contributing

We appreciate any form of contribution (feature request, bug report,
pull request, ...). We have no special requirements for Pull Request,
just follow the standard [GitHub way](https://help.github.com/articles/using-pull-requests/).


## License

The sources are release under a BSD 3-Clause License. The full terms of that
license can be found in `LICENSE` file of this repository.
