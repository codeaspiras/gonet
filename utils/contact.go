package utils

import (
	"bufio"
	"fmt"
	"os"
)

func contact(template string, args []interface{}) string {
	fmt.Printf(template, args...)
	buf := bufio.NewScanner(os.Stdin)
	if buf.Scan() {
		return buf.Text()
	}

	if err := buf.Err(); err != nil {
		panic(err.Error())
	}

	return ""
}
