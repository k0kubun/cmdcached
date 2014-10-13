# Cmdcached

Command result cache server for your development.  
You can execute any static commands faster by cmdcached.  
  
**WORK IN PROGRESS**

## Installation

```bash
$ go get github.com/k0kubun/cmdcached
```

## Example usage

First of all, you should start cmdcached server.

```bash
$ cmdcached
```

### 1. ghq list

Example usage for [ghq](https://github.com/motemen/ghq). You can executes `ghq list` faster by cmdcached.

#### Prefix cmdcached

```bash
$ cmdcached ghq list
```

You can use cached result by using `cmdcached ghq list` instead of `ghq list`.

#### Write ~/.cmdcached

To update cache on directory structure change, write ~/.cmdcached with [toml](https://github.com/toml-lang/toml).

```
[[cache]]
command = "ghq list"
subscribe = "/home/k0kubun/src"
```

### 2. git ls-files

If you work on git repository with 10000+ files, `git ls-files` will be slow. Cmdcached can make it faster.

#### Use cached result

```bash
$ cmdcached git ls-files
```

#### Write ~/.cmdcached

```
[[cache]]
command = "git ls-files"
each_directory = true
```

If you enables `each_directory` option, current directory will be used as cache key too.
Thus you can cache `git ls-files` per repository. The directory's file structure is subscribed.

## Pending features

I'm sorry but this project is **WORK IN PROGRESS**

- Core
  - file event subscription (`subscribe` directive)
    - recursive subscription
- Convenience
  - check server process availability
    - start server from client
