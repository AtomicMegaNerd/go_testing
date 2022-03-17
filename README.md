# Unit Testing in Go

This is some notes based on the Go Testing PluralSight.

## Go Standard library

Go has several packages in the standard library which are useful for learning how to write tests:

- testing
- testing/quick - Useful for black-box testing.
- testing/iotest - Designed for testing Reader and Writer interfaces.
- net/http/httptest - For testing HTTP connections and interfaces.

## Community Projects

### Testify

Testify is a library designed to support assertions in Go tests that works well with the standard
library:

[https://github.com/stretchr/testify](https://github.com/stretchr/testify)

### Ginkgo

This is for writing behavioural (BDD) style tests:

[https://github.com/onsi/ginkgo](https://github.com/onsi/ginkgo)

### GoConvey

Let's you use the browser to view test results:

[http://goconvey.co/](http://goconvey.co/)

### HTTP Expect

This is a project that is geared towards testing REST API's.

[https://github.com/gavv/httpexpect](https://github.com/gavv/httpexpect)

### Go Mock

This is a project for writing Mocks in Go:

[https://github.com/golang/mock](https://github.com/golang/mock)

### Go-SQLMock

Full SQL database provider that is an in-memory mockable database:

[https://github.com/golang/mock](https://github.com/golang/mock)

