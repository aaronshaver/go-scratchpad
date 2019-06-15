package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isDeranged(slice1, slice2 []string) bool {
	for i := range slice1 {
		if slice1[i] == slice2[i] {
			return false
		}
	}

	return true
}

func getInput() []string {
	reader := bufio.NewReader(os.Stdin)

	var items []string

	fmt.Println("Type a string, then press <ENTER>. Press <ENTER> by itself to stop.")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\r\n", "", -1) // trim line ending on Windows
		if len(text) == 0 {                          // exit for loop when user hits ENTER key
			break
		} else {
			items = append(items, text)
		}
	}

	return items
}

func main() {
	var items = getInput()
	fmt.Println(items)
}
