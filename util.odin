package main

import "core:fmt"
import "core:os"

fatalf :: proc(format: string, args: ..any) {
	fmt.printfln(format, ..args)
	os.exit(1)
}
