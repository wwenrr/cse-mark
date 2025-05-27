package main

import infrastructure "thuanle/cse-mark/internal/infra"

func main() {
	infrastructure.InitZerolog()
	_ = infrastructure.InitDotenv()

}
