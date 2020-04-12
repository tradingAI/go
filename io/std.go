package io

import "os"

// StdinNotEmpty ..
func StdinNotEmpty() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
