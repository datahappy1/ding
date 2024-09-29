package main

import (
	"flag"
	"os"

	"github.com/datahappy1/ding/internal"
)

const DEFAULT_HOST = "www.seznam.cz"
const MAX_ITERATIONS = 10
const TIME_TO_BEEP_DEFAULT_MULTIPLICATOR = 10

var (
	defaultHostFlag string
	maxItersFlag    int
)

func main() {
	flag.StringVar(&defaultHostFlag, "default_host", DEFAULT_HOST, "default host")
	flag.IntVar(&maxItersFlag, "max_iterations", MAX_ITERATIONS, "max iterations to proceed")
	flag.Parse()

	internal.Run(defaultHostFlag, maxItersFlag, TIME_TO_BEEP_DEFAULT_MULTIPLICATOR)
	os.Exit(0)
}
