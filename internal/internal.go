package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/gen2brain/beeep"
)

const DEFAULT_HOST = "www.seznam.cz"
const MAX_ITERATIONS = 50
const DEFAULT_TIME_MS = 100
const TIME_TO_BEEP_DEFAULT_MULTIPLICATOR = 10

func parse_time_from_line(line string) int {
	re := regexp.MustCompile(`time=\d*`)
	re_search_result := re.FindString(string(line))

	time_str := strings.Split(re_search_result, "=")

	time_value := 0
	var err error

	if len(time_str) > 1 {
		time_value, err = strconv.Atoi(time_str[1])
		if err != nil {
			panic(err)
		}
	} else {
		time_value = DEFAULT_TIME_MS
	}
	return time_value
}

func generate_beep(length int) {
	err := beeep.Beep(beeep.DefaultFreq, length)
	if err != nil {
		panic(err)
	}
}

func Run(host string, iterations int) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "INFO: ", log.Lshortfile)

		infof = func(info string) {
			logger.Output(2, info)
		}
	)

	ping_command := ""
	if host != "" {
		ping_command = fmt.Sprintf("ping %s -t ", host)
	} else {
		ping_command = fmt.Sprintf("ping %s -t ", DEFAULT_HOST)
	}

	cmdArgs := strings.Fields(ping_command)

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	stdout, _ := cmd.StdoutPipe()

	cmd.Start()
	oneByte := make([]byte, 100)
	num := 1

	for {
		_, err := stdout.Read(oneByte)
		if err != nil {
			panic(err)
		}
		r := bufio.NewReader(stdout)
		line, _, _ := r.ReadLine()
		infof(string(line))

		if strings.HasPrefix(string(line), "with 32 bytes of data:") {
			continue
		}

		time_value := parse_time_from_line(string(line))

		fmt.Println("ping round trip time (seconds):", time_value)

		generate_beep(time_value * TIME_TO_BEEP_DEFAULT_MULTIPLICATOR)

		num = num + 1

		if num > MAX_ITERATIONS {
			break
		}
	}
	infof("reached max iterations count")
	os.Exit(0)
}
