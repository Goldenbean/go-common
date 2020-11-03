package common

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func ReadFile(filePath string, queue chan<- string) {
	fin, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer fin.Close()

	scanner := bufio.NewScanner(fin)
	for scanner.Scan() {
		line := scanner.Text()
		queue <- line
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func WriteFile(filePath string, outQ <-chan []string) {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
	}
	// Write Unmarshaled json data to CSV file
	w := csv.NewWriter(f)
	valid := true
	for valid {
		record := <-outQ
		w.Write(record)
		w.Flush()
	}

}
