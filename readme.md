# gol

gol is a GO logging library with simple frame tracing.

### Installation

`go get github.com/rockwithamoon/gol`

### Usage

```go
package main

import (
    "errors"
    "github.com/rockwithamoon/gol"
)

func main() {
    // with printf style arguments
    gol.Info("Info %v", "message")

    // with errors
    err := errors.New("Error")
    gol.Error(err)

    // set custom level
    gol.SetLevel(gol.DEBUG)
    gol.Debug("Debug message")
}
```

Will output

```s
INFO: Info message (/Users/user/go/src/prog/main.go,:10 main.main)
ERROR: Error (/Users/user/go/src/prog/main.go,:14 main.main)
DEBUG: Debug message (/Users/user/go/src/prog/main.go,:18 main.main)
```

### Reference

`godoc` or [http://godoc.org/github.com/rockwithamoon/gol](http://godoc.org/github.com/rockwithamoon/gol)
