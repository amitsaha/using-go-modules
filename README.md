# Using go modules

Go switched to using modules by default starting with Go 1.16. This essentially means
that all Go packages must be a module.

## Zen of Go Modules

0. A *repository* can contain *one* or more *modules*. One module can contain one ore more *packages*
1. It's important that the path you specify to  `go get github.com/username/dirname/path` matches
what you have in the `go.mod` file i.e. the module name must be: `github.com/username/dirname/path`.
2. In your code, you import packages, not modules
3. When using private packages, use the GONOSUMDB variable as explained [here](https://medium.com/mabar/today-i-learned-fix-go-get-private-repository-return-error-reading-sum-golang-org-lookup-93058a058dd8)

## Examples

This repository contains examples of achieving common tasks when using modules. Each
directory is a Go module and has been created as follows:

```
mkdir <dir-name>
cd <dir-name>
go mod init github.com/username/<dir-name>
```
- [Simple hello world with sub-packages, but no external dependency](./hello-world)
- [Use an external package](./hello-dep)
- TODO: Using a changed/local version of an external package
