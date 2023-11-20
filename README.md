# XNS gateway
XNS gateway can support [XNS contract](https://github.com/Moonyongjung/xns-contract). (XNS contract is not public yet.)

## Introduction

### What is XNS?

Most addresses of blockchains are difficult to recognize as address is generated randomly by using key algorithm. Also, an address is too long. For easily reminding address, naming service is that short name (as 'domain') is connected to the account address of the blockchain similarly DNS. This contract supports to operate methods are creating/retrieving domain in the XNS system. It is inspired by [ENS](https://ens.domains/), but logic and architecture of the XNS are not the same.

### Gateway

XNS gateway is a program that supports easily XNS contract deployment for business owners and allows owners to conveniently operate the settings they need to make through CLI. In addition, it was developed to support XPLA's explorer or vault to communicate through HTTP REST API for contract inquiry to use XNS.

## Version compatibility

|XNS gateway|XNS contract|Note|
|:--:|:--:|:--:|
|v0.0.2|v0.2.0||
|v0.0.1|v0.1.0||

## Install

Generate go binary file.

```
$ make install
```

## Deploy XNS contracts

### Prerequisites
Copy XNS contracts to the directory of the XNS gateway. (`./contract/artifacts`)

### Run
```shell
# 1. Initialization
$ xns init

# 2. Recover key
$ xns key recover

# 3. Store contracts
$ xns execute contract store-all

# 4. Instantiate contracts
$ xns execute contract instantiate-all

# 5. Connects contracts
$ xns execute contract connect

# 6. Start gateway
$ xns start
```

## APIs

XNS gateway provides CLI by using binary file. The usage of each can be confirmed when running the binary file, and since each command interacts with the XNS contract, deploying of the contract must be preceded.

Also, check APIs in [swagger](./docs/swagger.md)