package main

import (
	"os"
	"time"

	clockface "github.com/tsobe/lgwt/maths"
)

func main() {
	t := time.Now()
	clockface.WriteSVG(os.Stdout, t)
}
