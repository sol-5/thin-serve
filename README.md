# Thin Server
A small, lightweight and portable server written in GoLang. This server can be used for serving folders and static websites on the local network.

# Why GoLang?
GoLang is an incredibly powerful yet minimalistic language. The entire code for thin server comes out to only 30 Lines (with comments!)

## Usage
```
Usage of thin-server:
  -folder string
    	Folder to serve on the server (default is current directory)

  -port int
    	System port for running the server (default 8080)
```

## Compilation
```bash
$ go build thin-server.go
```
## TODO
Rewrite in Rust or C++ to better understand Request-Response cycles, Sockets and Concurrency.
