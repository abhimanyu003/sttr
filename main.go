package main

import (
	"github.com/abhimanyu003/sttr/cmd"
)

// version specify version of application using ldflags
var version = "dev"

// generate the individual processor commands
//go:generate go run cmd/generate.go
func main() {
	cmd.Version = version
	cmd.Execute()
}
