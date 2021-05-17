package fakehttpclient

import (
	"errors"
	"net/http"
)

type roundTripHandler struct {
	context *roundTripperContext
}

type roundTripperContext struct {
	responses []*http.Response
	cursor    int
	maxRun    int
}

func (rtc *roundTripperContext) RoundTrip(_ *http.Request) (*http.Response, error) {
	if rtc.cursor >= rtc.maxRun {
		return nil, errors.New("client exhausted")
	}

	index := rtc.cursor % len(rtc.responses)
	rtc.cursor++

	if rtc.responses[index] == nil {
		return nil, errors.New("http: nil Request")
	}

	return rtc.responses[index], nil
}

func (rth roundTripHandler) RoundTrip(req *http.Request) (*http.Response, error) {
	return rth.context.RoundTrip(req)
}

func NewFakeHTTPClient(responses []*http.Response, maxRun int) (*http.Client, error) {
	if len(responses) == 0 {
		return nil, errors.New("empty list of responses")
	}

	if maxRun <= 0 {
		return nil, errors.New("non-positive maxRun")
	}

	client := &http.Client{}
	client.Transport = roundTripHandler{
		context: &roundTripperContext{
			responses: responses,
			cursor:    0,
			maxRun:    maxRun,
		},
	}
	return client, nil
}
