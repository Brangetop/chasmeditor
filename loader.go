package main

import "os"

func loadFileText(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err

		// add status bar interface later here
	}
	return string(b), nil
}
