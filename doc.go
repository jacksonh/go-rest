/*
## go-rest A minimalistic REST framework for Go

### Reflection, Go structs, and JSON marshalling FTW!

* go get github.com/ungerik/go-rest
* import "github.com/ungerik/go-rest"
* Documentation: http://go.pkgdoc.org/github.com/ungerik/go-rest
* License: Public Domain

Download, build and run example:

	go get github.com/ungerik/go-rest
	cd github.com/ungerik/go-rest/example
	go install && example

The framework consists of only three functions:
HandleGet, HandlePost, RunServer.

Discussion:

This package can be considered bad design because
HandleGet and HandlePost use dynamic typing to hide 36 combinations
of handler function types to make the interface _easy_ to use.
36 static functions would have been more lines of code but
dramatic _simpler_ in their individual implementations.
So simple in fact, that there wouldn't be a point in
abstracting them away in an extra framework.
See this great talk about easy vs. simple:
http://www.infoq.com/presentations/Simple-Made-Easy
Rob Pike may also dislike this approach:
https://groups.google.com/d/msg/golang-nuts/z4T_n4MHbXM/jT9PoYc6I1IJ
So this package may be seen as an anti-pattern to all
that is good and right about Go.
Why use it then? Well, it is _easy_: it brings the good
parts of dynamic behaviour to Go.
Yes, that introduces some internal complexity,
but this complexity is still very low in absolute terms
and thus easy to control and debug.
The complexity of the dynamic code also does not spill over
into the package users' code, because the package user
uses structs and all the static typed goodness that
come with them.

Now let's get started with this little dynamic madness,
maybe it's useful and fun after all:

HandleGet uses a handler function that returns a struct or string
to create the GET response. Structs will be marshalled as JSON,
strings will be used as body with auto-detected content type.

Format of GET handler:

	func([url.Values]) ([struct|*struct|string][, error]) {}

Example:

	type MyStruct struct {
		A in
		B string
	}

	rest.HandleGet("/data.json", func() *MyStruct {
		return &MyStruct{A: 1, B: "Hello World"}
	})

	rest.HandleGet("/index.html", func() string {
		return "<!doctype html><p>Hello World"
	})

The GET handler function can optionally accept an url.Values argument
and return an error as second result value that will be displayed as
500 internal server error if not nil.

Example:

	rest.HandleGet("/data.json", func(params url.Values) (string, error) {
		v := params.Get("value")
		if v == "" {
			return nil, errors.New("Expecting GET parameter 'value'")
		}
		return "value = " + v, nil
	})

HandlePost maps POST form data or a JSON document to a struct that is passed
to the handler function. An error result from handler will be displayed
as 500 internal server error message. An optional first string result
will be displayed as a 200 response body with auto-detected content type.

Format of POST handler:

	func([*struct|url.Values]) ([struct|*struct|string],[error]) {}

Example:

	rest.HandlePost("/change-data", func(data *MyStruct) (err error) {
		// save data
		return err
	})

Both HandleGet and HandlePost also accept one optional string argument.
In that case handler is interpreted as an object and the string argument
as the name of the handler-method of this object.

Exampe:

	rest.HandleGet("/method-call", myObject, "MethodName")
*/
package rest