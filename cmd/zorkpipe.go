package main

import (
	"bufio"
	"example.com/zorkpipe/internal/zork"
	"fmt"
	"os"
)

func main() {

	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// piped
		processPipe()
	}
	fmt.Printf("\n-----\n%s\n", zork.GetDescription())

	os.Exit(0)
}

func processPipe() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

}

/*
 * does the shortening of input returning a shortened string (if it can shorten it)
 */
// TODO : could the implementation/logic be improved?
func shorten(input string) string {
	if len(input) <= 2 {
		return input
	}

	firstChar := input[0:1]
	lastChar := input[len(input)-1:]
	abbrLen := len(input) - 2

	return fmt.Sprintf("%s%d%s", firstChar, abbrLen, lastChar)
}
