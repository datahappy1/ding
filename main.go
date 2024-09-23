package main

import (
	"github.com/datahappy1/ding/internal"
)

const DEFAULT_HOST = "www.seznam.cz"
const MAX_ITERATIONS = 50
const DEFAULT_TIME_MS = 100
const TIME_TO_BEEP_DEFAULT_MULTIPLICATOR = 10

func main() {
	internal.Run(DEFAULT_HOST, MAX_ITERATIONS)
}
