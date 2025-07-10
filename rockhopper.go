package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: rockhopper <url> [output_file]")
		os.Exit(1)
	}

	outputFile := ""
	if len(os.Args) > 2 {
		outputFile = os.Args[2]
	} else {
		outputFile = path.Base(os.Args[1])
	}

	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching URL: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: HTTP status code %d for URL %s\n", resp.StatusCode, url)
		os.Exit(1)
	}

	// Redirect standard output to the output file
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Creating File: %v\n", err)
		os.Exit(1)
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Printf("Error saving to file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Downloaded %s to %s\n", url, outputFile)
}
