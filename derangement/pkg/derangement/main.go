package main

import (
	"math/rand"
	"time"
)

// Tests for derangement. Only works with true sets (collection of distinct objects)
func isDeranged(slice1, slice2 []string) bool {
	for i := range slice1 {
		if slice1[i] == slice2[i] {
			return false
		}
	}

	return true
}

// Obviously this is brute force and not performant for large slices. This is a weekend
// project and I'm tired, so the CTCI-style perfect answer shall be a later exericse ;-)
func Derange(a []string) []string {
	tempStrings := append([]string(nil), a...)
	rand.Seed(time.Now().UnixNano())
	for {
		rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
		if isDeranged(tempStrings, a) {
			return a
		}
	}
}
