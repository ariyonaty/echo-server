# echo-server

A tiny HTTP echo server written in Go.
Returns request details and environment info as JSON.
Inspired/motivated for some quick testing when working with proxies.

Inspired by https://github.com/Ealenn/Echo-Server, but significantly smaller due to being written in Go as a static binary.

## Run locally

```bash
go run main.go
```

## Docker

Build the image:
```bash
docker image build -t echo-server .
```

Run the container:
```bash
docker container run --rm -it -p 8888:8888 echo-server:latest
```


## Verify
```json
❯ curl http://localhost:8888/ -H "X-Custom-Header: deadbeef" | jq
{
  "environment": {
    "HOME": "/",
    "HOSTNAME": "f55c90611bef",
    "PATH": "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
    "TERM": "xterm"
  },
  "host": {
    "hostname": "f55c90611bef",
    "ip": "192.168.1.123",
    "ips": []
  },
  "http": {
    "baseUrl": "",
    "method": "GET",
    "originalUrl": "/healthz",
    "protocol": "http"
  },
  "request": {
    "body": {},
    "cookies": {},
    "headers": {
      "Accept": "*/*",
      "User-Agent": "curl/8.7.1",
      "X-Custom-Header": "deadbeef"
    },
    "params": {
      "0": "/healthz"
    },
    "query": {}
  }
}
```
