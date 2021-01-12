package propaganda

import "net/http"

// The Input represents a middleware for a http.Handler.
//
//  server := http.Server{
//  	Handler: Input1(Input2(http.DefaultServeMux)),
//  }
//
type Input = func(http.Handler) http.Handler
