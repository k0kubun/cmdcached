# Cmdcached

Command result cache server for your development.  
You can execute any static commands faster by cmdcached.  

## Installation

```bash
$ go get github.com/k0kubun/cmdcached
```

## Example usage

Example usage for [ghq](https://github.com/motemen/ghq). You can executes `ghq list` faster by cmdcached.

### Write ~/.cmdcached

To update cache on directory structure change, write ~/.cmdcached with [toml](https://github.com/toml-lang/toml).

```
[[cache]]
command = "ghq list"
subscribe = "/home/k0kubun/src"
```

### Start server

You should start cmdcached server.

```bash
$ cmdcached
```

### Prefix cmdcached

```bash
$ cmdcached ghq list
```

You can use cached result by using `cmdcached ghq list` instead of `ghq list`.
