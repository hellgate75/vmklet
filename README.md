<p align="center" style="width: 100%"><img width="200" height="200" src="/images/golang.png" /></p>

# Go Virtual Machine Provisioning Scripting Parser (vmklet)

Go-Lang Virtual Machine provisioning parser language. This package allow to define provision machine via ssh, docker-compose, awsclient and so on ...


## Prerequisites

* [Go](https://golang.org/dl/) (tested with version 1.8)

## Goals

Define a virtual machine remote/local provisioner for OS, designed and optimized for [vmkube](https://github.com/hellgate75/vmkube).

## What is provided?

Provided features:

* Scripting language parse

* Scripting language helper

* Scripting language command

* Multi language translator (Ansible, kubecli, AWS CloudFormation, ....)


To compile and run this project you have to check availability of following software:
* [Go](https://golang.org/dl/) (tested with version 1.8)
* Test and Utility GOLang packages ([UUID Package](https://github.com/satori/go.uuid), [Unit Test](https://github.com/stretchr/testify)) and [GO SSH Terminal](http://golang.org/x/crypto/ssh/terminal), [YAML Parser](http://gopkg.in/yaml.v2)


## Architecture



## Configuration


## Checkout and test this repository

Go in you `GOPATH/src` folder and type :
```sh
 go get github.com/stretchr/testify
 go get github.com/satori/go.uuid
 go get golang.org/x/crypto/ssh/terminal
 go get gopkg.in/yaml.v2
 git clone https://github.com/hellgate75/vmklet.git

```
or simply :
```sh
 go get github.com/stretchr/testify
 go get github.com/satori/go.uuid
 go get golang.org/x/crypto/ssh/terminal
 go get gopkg.in/yaml.v2
 go get github.com/hellgate75/vmklet
```


## Build

It's present a make file that returns an help on the call :

```sh
make
```
Provided `Makefile` help returns following options :
```sh
make [all|init|test|build|exe|run|clean|install]
all: test build exe run
init: get required external packages
test: run unit test
build: build the module
exe: make executable for the module
clean: clean module C objects
run: exec the module code
install: install the module in go libs
```

Alternatively you can execute following commands :
 * `go get github.com/stretchr/testify` to download unit test external package
 * `go get github.com/satori/go.uuid` to download UUID management external package
 * `go get golang.org/x/crypto/ssh/terminal` to download SSH terminal external package
 * `go get gopkg.in/yaml.v2` to download YAML parser
 * `go build .` to build the project
 * `go test` to run unit and integration test on the project
 * `go run main.go` to execute the project
 * `go build --buildmode exe .` to create an executable command
 * `go build install` to install executable command


## Execution

The tool provides an help section, describing commands, sub-commands and has a nested help level for commands details.

The help is available executing : 
* `vmklet help` General Help
* `vmklet help [command]` Detailed Command syntax helper

Parser provides a sample for the expected input format. Import and Export of actions are provided in following file formats:
* JSON - standard JSON language
* XML - Untagged and un-described XML format (Pure XML tag sequence, no XML definition or version is accepted).
* YAML - standard YAML format.

In this release the command list is composed by following keys :

*NOT* *PROVIDED* *YET*

## License

Licensed under the [MIT](/LICENSE) License (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[https://opensource.org/licenses/MIT](https://opensource.org/licenses/MIT)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
