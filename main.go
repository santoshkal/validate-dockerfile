package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isMultiStage(lines []string) bool {
	fromCount := 0
	for _, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "FROM ") {
			fromCount++
		}
	}
	return fromCount > 1
}

func usesBaseImage(lines []string, baseImage string) bool {
	lastStageIndex := 0
	for idx, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "FROM ") {
			lastStageIndex = idx
		}
	}
	return strings.Contains(strings.TrimSpace(lines[lastStageIndex]), baseImage)
}

func hasInlineComments(lines []string) bool {
	for _, line := range lines {
		if !strings.Contains(line, "#") {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <path-to-dockerfile>")
		return
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if isMultiStage(lines) {
		fmt.Println("Dockerfile is multi-stage.")
	} else {
		fmt.Println("Dockerfile is not multi-stage.")
	}

	if usesBaseImage(lines, "scratch") {
		fmt.Println("Dockerfile uses scratch as the final base image.")
	}

	if usesBaseImage(lines, "alpine") {
		fmt.Println("Dockerfile uses alpine as the final base image.")
	}

	// Add more checks for other base images like wolfi-linux, golang, node, python, R, JavaScript, Bun, Rust, C, and Lua

	if hasInlineComments(lines) {
		fmt.Println("Dockerfile has inline comments for each line.")
	} else {
		fmt.Println("Dockerfile is missing inline comments for some lines.")
	}
}
