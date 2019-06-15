// This package accepts a set of strings as user input on the console, then
// creates a random shuffling of them such that the output is a derangement
// of the original set. This means no string in the output set is in the same
// position it was as in the original set.
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Tests for derangement. Only works well with true sets (collection of distinct objects)
func isDeranged(a, b []string) bool {
	for i := range a {
		if a[i] == b[i] {
			return false
		}
	}

	return true
}

// Obviously this is brute force and not performant for large slices. This is a weekend
// project and I'm tired, so the CTCI-style perfect answer shall be a later exericse ;-)
func derange(a []string) []string {
	tempStrings := append([]string(nil), a...)
	rand.Seed(time.Now().UnixNano())
	for {
		rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
		if isDeranged(tempStrings, a) {
			return a
		}
	}
}

func getInput() []string {
	reader := bufio.NewReader(os.Stdin)

	var items []string

	fmt.Println("Type a string, then press <ENTER>. Press <ENTER> by itself to stop.")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// DISCLAIMER: Windows-only; could do GOOS detection here to make it work on other OSes, but meh
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
	derange(items)
	fmt.Println(items)
}
