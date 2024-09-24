package main

import (
	"flag"
	"os"

	"github.com/datahappy1/ding/internal"
)

const DEFAULT_HOST = "www.seznam.cz"
const MAX_ITERATIONS = 10
const DEFAULT_TIME_MS = 100
const TIME_TO_BEEP_DEFAULT_MULTIPLICATOR = 10

func main() {
	var defaultHost = flag.String("default_host", DEFAULT_HOST, "default host")
	var maxIters = flag.Int("max_iterations", MAX_ITERATIONS, "max iterations to proceed")

	internal.Run(*defaultHost, *maxIters, DEFAULT_TIME_MS, TIME_TO_BEEP_DEFAULT_MULTIPLICATOR)
	os.Exit(0)
}
