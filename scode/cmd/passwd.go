package cmd

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

//password
var PwdKey = make([]byte, 9)

// compare slice
func ByteSliceEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// readPass
func readPass() []byte {
	fmt.Print("ENTER PASSWORD: ")
	bytepw, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		os.Exit(1)
	}
	return bytepw
}
