# GoDirBrowser

File sharing (download) application over HTTP, written in Go.

## How to use it

### Installing

Make sure that Go (tested on `go version 1.19 linux/amd64`) is installed on your computer.
```sh
./scripts/build.sh
./build/godirbrowser-linux-amd64 # The server is running on localhost:8080 by default.
```

## TODO

+ [ ] Possibility to use HTTPS
  + [ ] force https (http redirect)
+ [ ] POST requests
+ [ ] change the directory to serve on the command line
+ [ ] change IP:port with flags