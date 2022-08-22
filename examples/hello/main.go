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
