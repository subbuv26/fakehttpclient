# Fake HTTP Client

A fake http client that acts a http client and serves requests as per the needs of tests.

To test code that handles different kinds of http responses, this fake client comes handy.
As the client gets configured with the desired responses, the code that gets tested receives the same responses in the same order.

## Installation

```
go get github.com/subbuv26/fakehttpclient
```
