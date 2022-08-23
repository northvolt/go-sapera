# Go wrapper for Sapera LT SDK

Go language wrapper for the Sapera LT SDK.

https://www.teledynedalsa.com/en/products/imaging/vision-software/sapera-lt/

This package currently only works on Linux platforms.

To use it you must obtain a copy of the SaperaLT SDK for Linux from Teledyne DALSA directly.

## How To Use

```go
package main

import (
	"fmt"

	"github.com/northvolt/go-sapera"
)

func main() {
	fmt.Println("Initializing SaperaLT...")
	sapera.Init()

	fmt.Sprintln("Running go-sapera %s with Sapera SDK %s", sapera.Version(), sapera.SDKVersion())
	fmt.Sprintln("Detected %d boards", sapera.GetServerCount())
}
```

For an example of how to capture frame data, please see [examples/capture/main.go](./examples/capture/main.go)

## Installing

To use it you must obtain a copy of the Sapera LT SDK for Linux from Teledyne DALSA directly. See:

https://www.teledynedalsa.com/en/products/imaging/vision-software/sapera-lt/download/

The drivers for the board that you want to use must be installed on the Linux host operating system.

## Why it exists

Computer vision applications written in Go can easily take advantage of multi-core concurrency while also having access to packages such as GoCV (https://gocv.io). The `go-sapera` package now makes is possible for Go programs to connect to cameras and line scanners that are generally used for industrial applications.
