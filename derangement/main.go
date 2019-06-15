// This package accepts a set of strings as user input on the console, then
// creates a random shuffling of them such that the output is a derangement
// of the original set. This means no string in the output set is in the same
// position it was as in the original set.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aaronshaver/go-scratchpad/pkg/derangement"
)

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
	derangement.Derange(items)
	fmt.Println(items)
}
