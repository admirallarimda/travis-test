package main

import "log"
import "runtime"

func main() {
	log.Printf("Hello, %s", runtime.GOOS)
}
