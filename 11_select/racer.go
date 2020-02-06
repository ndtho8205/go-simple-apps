package main

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeout = 10 * time.Second

func Racer(firstURL, secondURL string) (string, error) {
	return ConfigurableRacer(firstURL, secondURL, tenSecondTimeout)
}

func ConfigurableRacer(firstURL, secondURL string, timeout time.Duration) (string, error) {
	select {
	case <-ping(firstURL):
		return firstURL, nil
	case <-ping(secondURL):
		return secondURL, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out")
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}

func Racer_v1(firstURL, secondURL string) string {
	firstDuration := measureResponseTime(firstURL)
	secondDuration := measureResponseTime(secondURL)

	if firstDuration < secondDuration {
		return firstURL
	}

	return secondURL
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()

	http.Get(url)

	return time.Since(start)
}
