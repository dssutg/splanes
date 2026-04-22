package main

import (
	"fmt"
	"os"
)

func fatalf(format string, args ...any) {
	fmt.Printf(format+"\n", args...)
	os.Exit(1)
}