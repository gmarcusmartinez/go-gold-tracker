package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"

	"fyne.io/fyne/v2/test"
	"github.com/gmarcusmartinez/gold-tracker/repository"
)

func TestMain(m *testing.M) {
	a := test.NewApp()

	testApp.App = a
	testApp.HTTPClient = client
	testApp.MainWindow = a.NewWindow("")

	testApp.DB = repository.NewTestRepository()

	os.Exit(m.Run())
}

var testApp Config

var client = NewTestClient(func(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
		Header:     make(http.Header),
	}
})

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
