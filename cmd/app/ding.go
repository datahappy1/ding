package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	host := flag.String("host", "", "host")
	iterations_count, err := strconv.Atoi(*flag.String("iterations", "", "iterations count"))
	if err != nil {
		panic(err)
	}

	flag.Parse()

	fmt.Println(ding.run(*host, iterations_count))

}
