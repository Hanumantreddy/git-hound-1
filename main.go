package main

import (
	_ "net/http/pprof"

	"github.com/hanumanthreddy/git-hound/cmd"
)

func main() {
	cmd.Execute()
}
