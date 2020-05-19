package main

import (
	"log"
	"os"

	_ "gobook/chapter2/sample/matchers"
	"gobook/chapter2/sample/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
