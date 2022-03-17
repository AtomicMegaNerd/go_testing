# Unit Testing in Go

This is some notes based on the Go Testing PluralSight.

---

## Go Standard library

Go has several packages in the standard library which are useful for learning how to write tests:

- testing
- testing/quick - Useful for black-box testing.
- testing/iotest - Designed for testing Reader and Writer interfaces.
- net/http/httptest - For testing HTTP connections and interfaces.

---

## Go Testing Naming Conventions

- Always add a \_test suffix to test files in Go. These files are excluded from release binaries!
- Always name functions TestFunctionName.
- Accept one parameter - \*testing.T
- Use the same package name as the file you are testing if you want "white box" tests with access
  to internal resources.
- Alternatively, use package\_test for "black box" tests where you do not have access to internal
  package private resources.

## Go Test Command

```bash
go test # Run all tests in current directory
```

```bash
go test -run {pkg1} {pkg2} ... {pkgn} # Test in specific packages
```

```bash
go test ./... # Run tests in current package and descendants
```

```bash
go test -v # Generate verbose output
```

```bash
go test -run {regexp} # Run any tests matching {regexp}
```

```bash
go test -cover # Shows code coverage total
```

```bash
go test -coverprofile cover.out # Generate code coverage report to file cover.out
```

```bash
go tool cover -func cover.out # Parse coverage report and show coverage totals per function
```

```bash
go tool cover -html cover.out # Parse coverage report and show coverage info in web browser
```

---

## Reporting Test Failures

### Immediate Failure

Will exit the failing test immediately. It will exit the current test function but not the test
suites.

```go
t.FailNow()
```

```go
t.Fatal(args... interface{})
```

```go
t.Fatalf(format string, args... interface{})
```

### Non-Immediate Failure

The test will be marked as failed but the current test function will keep running. This is useful
if you want to test multiple things about an object in a given test and make sure all of them
do run even if one or more conditions do fail.

```go
t.Fail()
```

```go
t.Error(args... interface{})
```

```go
t.Errorf(format string, args... interface{})
```

---

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

