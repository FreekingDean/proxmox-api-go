package main

import (
	"fmt"
	"io"
	"net/http"
)

func download() (string, error) {
	resp, err := http.Get(apijs)
	if err != nil {
		return "", fmt.Errorf("download err: %w", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read err: %w", err)
	}
	return string(body), nil
}
