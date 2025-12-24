/*
Package head is an HTTP middleware for Go that facilitates adding of X-Original-Uri.

	package main

	import (
	    "net/http"

	    "github.com/nikiminoj/head"
	)

	var myHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	    w.Write([]byte("hello world"))
	})

	func main() {
	    headMiddleware := head.New(head.Options{})

	    app := headMiddleware.Handler(myHandler)
	    http.ListenAndServe("127.0.0.1:3000", app)
	}
*/
package head
