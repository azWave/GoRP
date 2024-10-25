package game

import (
	"bufio"
	"strings"
	"os"
)

func ReadInput() []string {
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    return strings.Split(strings.TrimSpace(input), "")
}