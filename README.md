[![GoDoc](https://godoc.org/github.com/subbuv26/fakehttpclient?status.svg)](https://pkg.go.dev/github.com/subbuv26/fakehttpclient)
[![Go Report Card](https://goreportcard.com/badge/github.com/subbuv26/fakehttpclient)](https://goreportcard.com/report/github.com/subbuv26/fakehttpclient)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

***Archived***
Please visit https://github.com/subbuv26/mockhttpclient

# Fake HTTP Client

A fake http client that acts a http client and serves requests as per the needs of tests.

To test code that handles different kinds of http responses, this fake client comes handy.
As the client gets configured with the desired responses, the code that gets tested receives the same responses in the same order.

## Installation

```
go get github.com/subbuv26/fakehttpclient
```

## Usage
### Example 1:
The function to be tested (createAndVerifyResource) has the below functionality
1. CREATE resource using POST call, which returns http OK
2. GET the resource using GET call, which returns http Service Unavailable (Server Busy)
3. retry to GET the resource, which returns http OK
4. On Success return true, otherwise false


```go

import fakehc "github.com/subbuv26/fakehttpclient"

resp1 := &http.Response{
    StatusCode: 200,
    Header:     http.Header{},
    Body:       ioutil.NopCloser(bytes.NewReader([]byte("body"))),
}

resp2 := &http.Response{
    StatusCode: 503,
    Header:     http.Header{},
    Body:       ioutil.NopCloser(bytes.NewReader([]byte("body"))),
}

responseMap := make(fakehc.ResponseConfigMap)

responseMap[http.MethodPost] = &ResponseConfig{}
responseMap[http.MethodPost].Responses = []*http.Response{resp1}
responseMap[http.MethodGet] = &ResponseConfig{}
responseMap[http.MethodGet].Responses = []*http.Response{resp2, resp1}

// Create Client
client, _ := NewFakeHTTPClient(responseMap)

// myRESTClient is the client that is going to be tested
myRESTClient.client = client

// createAndVerifyResource creates a resource  with POST and verifies with GET
// with myRESTClient.client
ok := myRESTClient.createAndVerifyResource(Resource{})
if ok {
	// Success 
} else {
	// Failed
}

```

### Example 2
The function to be tested (createResources) creates N number of resources by calling continuous POST calls

```go
import fakehc "github.com/subbuv26/fakehttpclient"

N := 5

resp1 := &http.Response{
    StatusCode: 200,
    Header:     http.Header{},
    Body:       ioutil.NopCloser(bytes.NewReader([]byte("body"))),
}

responseMap := make(fakehc.ResponseConfigMap)

responseMap[http.MethodPost] = &ResponseConfig{}
responseMap[http.MethodPost].Responses = []*http.Response{resp1}
responseMap[http.MethodPost].MaxRun = N

// Create Client
client, _ := NewFakeHTTPClient(responseMap)

// myRESTClient is the client that is going to be tested
myRESTClient.client = client

// createResources makes post calls using myRESTClient.client
ok := myRESTClient.createResources([]Resources)
if ok {
    // Success 
} else {
    // Failed
}

```



