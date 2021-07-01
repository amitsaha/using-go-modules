# Using go modules

Go switched to using modules by default starting with Go 1.16. This essentially means
that all Go packages must be a module.

## Examples

This repository contains exammples of achieving common tasks when using modules. Each
directory is a Go module and has been created as follows:

```
mkdir <dir-name>
cd <dir-name>
go mod init github.com/username/<dir-name>
```
- [Simple hello world with sub-packages, but no external dependency](./hello-world)
- [Use an external package](./hello-dep)
- TODO: Using a changed/local version of an external package

