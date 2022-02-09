package main

import (
	"fmt"
	"net/http"
	"time"
)

type PingSignal chan interface{}

func Racer(urlA string, urlB string) (string, error) {
	return ConfigurableRacer(urlA, urlB, 10*time.Second)
}

func ConfigurableRacer(urlA string, urlB string, timeout time.Duration) (string, error) {
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out in %s", timeout)
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
