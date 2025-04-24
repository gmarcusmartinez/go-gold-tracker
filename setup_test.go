package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

var testApp Config

var jsonToReturn = `
	{
		"ts": 1745479501939,
		"tsj": 1745479495567,
		"date": "Apr 24th 2025, 03:24:55 am NY",
		"items": [
			{
				"curr": "USD",
				"xauPrice": 3331.895,
				"xagPrice": 33.3823,
				"chgXau": 46.5375,
				"chgXag": -0.1854,
				"pcXau": 1.4165,
				"pcXag": -0.5523,
				"xauClose": 3285.3575,
				"xagClose": 33.56773
			}
		]
	}
`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}
