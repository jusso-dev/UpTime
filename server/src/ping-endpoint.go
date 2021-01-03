package main

import (
	"crypto/tls"
	"net/http"
	"net/http/httptrace"
	"time"
)

//RESPONSE provides return object for pingURL func
type RESPONSE struct {
	SUCCESS bool  `json:"success" binding:"required"`
	TIME    int64 `json:"time" binding:"required"`
}

func timeGet(url string) RESPONSE {

	req, _ := http.NewRequest("GET", url, nil)
	var res RESPONSE

	var start, connect, dns, tlsHandshake time.Time

	trace := &httptrace.ClientTrace{
		DNSStart: func(dsi httptrace.DNSStartInfo) { dns = time.Now() },
		DNSDone: func(ddi httptrace.DNSDoneInfo) {
		},

		TLSHandshakeStart: func() { tlsHandshake = time.Now() },
		TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
		},

		ConnectStart: func(network, addr string) { connect = time.Now() },
		ConnectDone: func(network, addr string, err error) {
		},

		GotFirstResponseByte: func() {
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	start = time.Now()
	if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
		res.SUCCESS = false
		res.TIME = 0

		return res
	}

	res.SUCCESS = true
	res.TIME = time.Since(start).Milliseconds()

	return res
}

/*
	Function to ping supplied URL
*/
func pingURL(endpoint string) RESPONSE {
	res := timeGet(endpoint)
	return res
}
