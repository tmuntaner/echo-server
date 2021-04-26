# Test Echo Server

This is a simple web server which returns the request information back to the response.

## Example

```bash
curl -H "DNT: 1" -H "Authorization: Basic YWRtaW46c2VjcmV0" http://localhost:8080/echo                                                                                                                                                                                           Mon 26  9:35AM
```

```text
URL: /echo
Method: GET
Protocol: HTTP/1.1

Headers:
Accept: */*
Authorization: Basic YWRtaW46c2VjcmV0
Dnt: 1
User-Agent: curl/7.76.0
```

## Build and Run with Docker

```bash
docker build -t echo-server .
docker run --rm -p 8080:8080 echo-server
```