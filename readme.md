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
    defer gol.Debug("Debug message")

    // with printf style arguments
    gol.Info("Info %v", "message")

    // with errors
    err := errors.New("Error")
    gol.Error(err)

    // set custom level
    gol.SetLevel(gol.DEBUG)
}
```

Will output

```s
INFO: Info message (/Users/user/go/src/x/main.go,:12 main.main)
ERROR: Error (/Users/user/go/src/x/main.go,:16 main.main)
DEBUG: Debug message (/Users/user/go/src/x/main.go,:20 main.main)
```

### Reference

`godoc` or [http://godoc.org/github.com/rockwithamoon/gol](http://godoc.org/github.com/rockwithamoon/gol)
