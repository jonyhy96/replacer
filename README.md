# replacer
[![Go Report Card](https://goreportcard.com/badge/github.com/jonyhy96/replacer)](https://goreportcard.com/report/github.com/jonyhy96/replacer)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

replacer helps you to batch replace files

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

```
$ go get github.com/jonyhy96/replacer
$ replacer -h
$ replacer -f ./keys.json -w .
```

### Prerequisites

For develop

 - go ^1.10

### Installing

```
$ go build -o replacer .
$ sudo mv replacer /usr/local/bin
```

### Params

| param | required | e.g. |
| :--------: | :-----: | :----: |
| f     | true | -f ./keys.json |
| w     | true | -w .           |
| e     | false| -e go.mod,go.sum |

### Coding style

[CODEFMT](https://github.com/golang/go/wiki/CodeReviewComments)

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://gitlab.domain.com/golang/containerM/tags). 

## Authors

* **HAO YUN** - *Initial work* - [haoyun](https://github.com/jonyhy96)

See also the list of [contributors](CONTRIBUTORS.md) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

* nothing
