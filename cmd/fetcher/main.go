package main

import (
	"thuanle/cse-mark/internal"
	"thuanle/cse-mark/internal/services/fetcher"
)

func main() {
	internal.Load()

	fetcher.Execute()

	defer internal.Unload()
}
