# ghq-listd

Server-client based fast `ghq list`.  
This only works as an replacement for [ghq](https://github.com/motemen/ghq)'s `list` command.

## Why ghq-listd?

Whenever you execute `ghq list`, it causes disk I/O.  
Thus sometimes the execution will be very slow if you use HDD.  
  
So this ghq list daemon **provides a repository list cached on memory**.  
It is updated by file system notification on your ghq directory.  

## Installation

```bash
$ go get github.com/k0kubun/ghq-listd
```

## Usage

### Simple way

Just replace your `ghq list` with:

```bash
$ ghq-listd
```

On the first execution, it will start ghq-listd server automatically.  
And then ghq-listd client will interact with ghq-listd server through unix domain socket.  
From second execution, it will be fast.  

### Explicit way

```bash
$ ghq-listd server
```

It ensures that ghq-listd server is started.  
You can write this in your `.bashrc`.

```bash
$ ghq-listd client
```

This is subset of `ghq-listd`.  
It does not check ghq-listd server is alive or not.

## Acknowledgements

I have no discontent for ghq on SSD.  
This product is designed to be used on HDD environment.  
  
`ghq-listd` provides only `ghq list` function.  
Even if you use this product, you still need to use original ghq for other sub-commands.
