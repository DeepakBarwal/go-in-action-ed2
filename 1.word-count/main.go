package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("need to provide atleast one filename!")
		os.Exit(1)
	}

	// Loop over all arguments from index 1 onwards.
	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)

		// If there's an error, print it and skip to the next file.
		if err != nil {
			log.Printf("%s: %v", filename, err)
			continue
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		var wordCount int

		for scanner.Scan() {
			wordCount++
		}

		if scanner.Err() != nil {
			log.Printf("scan error %s: %v", filename, scanner.Err())
			file.Close()
			continue
		}

		file.Close()
		fmt.Printf("%s: %d\n", filename, wordCount)
	}
}
