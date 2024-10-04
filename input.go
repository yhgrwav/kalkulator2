package input

import (
	"bufio"
	"os"
	"strings"
)

// ReadInput считывает ввод пользователя
func ReadInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}
