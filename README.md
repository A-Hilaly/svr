# svr 

[![CircleCI](https://circleci.com/gh/A-Hilaly/svr.svg?style=svg)](https://circleci.com/gh/A-Hilaly/svr)

---

Command line to run simple servers for developpement purposes
### features

- http server responding to all patterns /*
- http to https redirection server for all patterns
- redicrection server for a specific pattern to unique url

### Getting started

```bash
go get -u github.com/a-hilaly/svr/...
```

### Usage

```bash
A collection of simple HTTP servers
                - Simple HTTP server
                - Simple HTTP proxy server
                - HTTP to HTTPS proxy server

Usage:
  svr [flags]

Flags:
  -h, --help              help for svr
  -s, --server string     Server type (default "default")
  -n, --pattern string    pattern url (default "/")
  -p, --port int          Server listening port (default 8080)
  -r, --redirect string   pattern url (default "http://www.google.com")
```

## Examples

##### simple server
```bash
$ svr -p 9000
2018/12/29 17:39:39 Simple server
2018/12/29 17:39:39 Listening on port 9000 (pattern: "/")

$ curl localhost:9000
``` 

##### redirection server

```bash
$ svr -s redirect -n /google -r https://google.com/
2018/12/29 17:41:05 Proxy server
2018/12/29 17:41:05 Listening on port 8000 (pattern: "/google")

$ curl localhost:8000
```

##### http to https

```
$ svr http2https -p 8000
2018/12/29 17:43:05 Http to https
2018/12/29 17:43:05 Listening on port 8080 (pattern: "/")
```

### docker images:

[hub.docker.io/r/ahilaly](https://cloud.docker.com/repository/docker/ahilaly/svr)