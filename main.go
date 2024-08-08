package main

import (
	"bufio"
	"daily_matter/logic/command"
	"daily_matter/logic/dailymatter"
	"daily_matter/logic/mattermanager"
	"os"
	"strings"
)

func main() {
	mattermanager.Init()
	dailymatter.Init()

	dp := dailymatter.DisplayConsolePacker{}
	cm := command.NewCommandManager(dp)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		args := strings.Split(input.Text(), " ")
		cm.Manager(args...)

		if input.Text() == "exit" {
			break
		}
	}
}
