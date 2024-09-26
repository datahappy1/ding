package main

import (
	"flag"
	"os"

	"github.com/datahappy1/ding/internal"
)

const DEFAULT_HOST = "www.seznam.cz"
const MAX_ITERATIONS = 10
const DEFAULT_TIME_MS = 10
const TIME_TO_BEEP_DEFAULT_MULTIPLICATOR = 30

var (
	defaultHostFlag string
	maxItersFlag    int
)

func main() {
	flag.StringVar(&defaultHostFlag, "default_host", DEFAULT_HOST, "default host")
	flag.IntVar(&maxItersFlag, "max_iterations", MAX_ITERATIONS, "max iterations to proceed")
	flag.Parse()

	internal.Run(defaultHostFlag, maxItersFlag, DEFAULT_TIME_MS, TIME_TO_BEEP_DEFAULT_MULTIPLICATOR)
	os.Exit(0)
}
