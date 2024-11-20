package main

import (
	"github.com/goyek/x/boot"
	"github.com/wasilibs/tools/tasks"
)

func main() {
	tasks.Define(tasks.Params{
		LibraryName: "shellcheck",
		LibraryRepo: "koalaman/shellcheck",
		GoReleaser:  true,
	})
	boot.Main()
}
