package main

import (
	"thuanle/cse-mark/internal"
	"thuanle/cse-mark/internal/services/tele"
)

func main() {
	internal.Load()

	tele.Execute()

	defer internal.Unload()
}
