package internal

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/gen2brain/beeep"
)

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
		time_value = 0
	}
	return time_value
}

func generate_beep(length int) {
	err := beeep.Beep(beeep.DefaultFreq, length)
	if err != nil {
		panic(err)
	}
}

func Run(host string, max_iterations int, time_to_beep_default_multiplicator int) {

	fmt.Printf("running with host:%s, max iters:%d, default multiplicator:%d \n",
		host, max_iterations, time_to_beep_default_multiplicator)

	ping_command := fmt.Sprintf("ping %s -t ", host)

	cmdArgs := strings.Fields(ping_command)

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	stdout, _ := cmd.StdoutPipe()

	cmd.Start()
	oneByte := make([]byte, 100)
	num := 1

	for {
		_, err := stdout.Read(oneByte)
		if err != nil {
			fmt.Println("unexpected response, check the host name")
			return
		}
		r := bufio.NewReader(stdout)
		line, _, _ := r.ReadLine()
		// fmt.Println(string(line))

		if strings.HasPrefix(string(line), "with 32 bytes of data:") {
			continue
		}

		time_value := parse_time_from_line(string(line))

		fmt.Println("ping round trip time (seconds):", time_value)

		generate_beep(time_value * time_to_beep_default_multiplicator)

		num = num + 1

		if num > max_iterations {
			break
		}
	}
	fmt.Printf("reached max iterations count (%d)", max_iterations)

}
