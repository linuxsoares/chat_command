package main

import (
	"os"
	"strconv"
)

var CHAT_COMMAND_OPEN_AI_TOKEN string

var MAXDIFFSIZE int

func init() {
	CHAT_COMMAND_OPEN_AI_TOKEN = os.Getenv("OPENAI_API_KEY")
	MAXDIFFSIZE, err := strconv.Atoi(os.Getenv("MAXDIFFSIZE"))
	if err != nil {
		MAXDIFFSIZE = 1000 // default value
	}
	_ = MAXDIFFSIZE
}
