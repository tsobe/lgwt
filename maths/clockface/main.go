package main

import (
	"os"
	"time"

	clockface "lgwt/maths"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
