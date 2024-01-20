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
	} else {
		// from terminal
		if len(os.Args) != 2 {
			fmt.Printf("zork : puts zork goodness at the end of piped content\n")
			fmt.Printf("usage: <command with output> | zork\n")
			fmt.Printf("v0.0.1\n")
			os.Exit(1)
		}

		output := shorten(os.Args[1])
		fmt.Printf("%s\n", output)
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
