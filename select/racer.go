package main

import (
	"fmt"
	"net/http"
	"time"
)

type PingSignal chan interface{}

func Racer(urlA string, urlB string) (winner string, err error) {
	select {
	case <-ping(urlA):
		return urlA, err
	case <-ping(urlB):
		return urlB, err
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out in %s", 10*time.Second)
	}
}

func ping(url string) PingSignal {
	signal := make(PingSignal)
	go func() {
		_, _ = http.Get(url)
		close(signal)
	}()
	return signal
}
