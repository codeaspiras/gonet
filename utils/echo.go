package utils

import "fmt"

func Echo(text string, args ...interface{}) {
	fmt.Printf("# "+text+"\n", args...)
}
