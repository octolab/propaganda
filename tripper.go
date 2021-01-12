package propaganda

import "net/http"

// The Output represents a middleware for a http.RoundTripper.
//
//  client := http.Client{
//  	Transport: Output1(Output2(http.DefaultTransport)),
//  }
//
type Output = func(http.RoundTripper) http.RoundTripper
