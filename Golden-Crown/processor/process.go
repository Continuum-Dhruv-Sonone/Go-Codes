package processor

import (
	"bufio"
	"io"
	"strings"
)

// ProcessInput gets the inputs from the process
func ProcessInput(f io.Reader) (inputs [][]string, err error) {

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		inputs = append(inputs, strings.Split(scanner.Text(), " "))
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return
}
