package main

import (
	"os"
	"time"

	clockface "lgwt/maths"
)

func main() {
	t := time.Now()
	clockface.WriteSVG(os.Stdout, t)
}
