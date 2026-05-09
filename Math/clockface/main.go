package main

import (
	"os"
	"time"

	clockface "example.com/go/Math"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
