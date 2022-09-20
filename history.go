package main

import (
	"fmt"
	"os"
	"strings"
)

type History []string

const HISTORY_LIMIT = 10

const HISTORY_FILENAME = "shell_history"

func (h History) display() {
	for index, command := range h {
		fmt.Println(index, command)
	}
}

func (h History) toString() string {
	history := []string(h)
	return strings.Join(history, ",	")
}

func (h History) saveHistoryToFile(command string) {
	bs := []byte(h.toString())
	os.WriteFile(HISTORY_FILENAME, bs, 0666)
}

func (historyPointer *History) update(command string) {
	if len(*historyPointer) == HISTORY_LIMIT {
		*historyPointer = (*historyPointer)[1:]
	} else {
		*historyPointer = append(*historyPointer, command)
	}
}

func restoreHistory() History {

	byteSlice, err := os.ReadFile(HISTORY_FILENAME)
	if err != nil {
		fmt.Println("Error is in read file", err)
		//os.Exit(1)
	}
	return History(strings.Split(string(byteSlice), ","))
}
