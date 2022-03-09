# Leofy Go Packages

This directory contains packages for interacting with the Leofy
smart contracts from a Go programming environment.

# Package Guides

- `contracts`: Contains functions to generate the text of the contract code for the contracts in the `/contracts` directory.

- `templates`: Contains functions to return transaction templates
for common transactions and scripts for interacting with the Leofy smart contracts.

- `test`: Contains automated go tests for testing the functionality of Leofy smart contracts.

# Running tests
Go should be installed on the machine. [Instructions here](https://github.com/golang/go#download-and-install).

Navigate to /lib/go folder and lauch:
```go
    make test 
```
or inside /lib/go/test
```go
    go test 
```