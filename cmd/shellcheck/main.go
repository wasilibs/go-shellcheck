package main

import (
	"os"

	"github.com/wasilibs/go-shellcheck/internal/runner"
	"github.com/wasilibs/go-shellcheck/internal/wasm"
)

func main() {
	os.Exit(runner.Run("shellcheck", os.Args[1:], wasm.Shellcheck, os.Stdin, os.Stdout, os.Stderr, "."))
}
