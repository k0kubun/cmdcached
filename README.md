# Cmdcached

High performance command cache server.  
You can execute any static commands faster by cmdcached.

## Installation

```bash
$ go get github.com/k0kubun/cmdcached
```

## Example usage

### 1. ghq list

Example usage for [ghq](https://github.com/motemen/ghq). You can executes `ghq list` faster by cmdcached.

#### Prefix cmdcached

```bash
$ cmdcached ghq list
```

You can use cached result by using `cmdcached ghq list` instead of `ghq list`.
In the first execution, cmdcached server will be started.
You can explicitly start it by just executing `cmdcached`.

#### Write ~/.cmdcached

```
[ghq list]
  subscribe=/home/k0kubun/src
```

When `/home/k0kubun/src`'s directory structure changes,
command cache will be updated.

### 2. git ls-files

If you work on git repository with 10000+ files, `git ls-files` will be slow. Cmdcached can make it faster.

#### Use cached result

```bash
$ cmdcached git ls-files
```

#### Write ~/.cmdcached

```
[git ls-files]
  each_directory=true
```

If you enables `each_directory` option, current directory will be used as cache key too.
Thus you can cache `git ls-files` per repository. The directory's file structure is subscribed.
