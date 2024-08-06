package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	var files []string

	fortuneCommand := exec.Command("fortune", "-f")
	pipe, err := fortuneCommand.StderrPipe()
	if err != nil {
		panic(err)
	}

	fortuneCommand.Start()
	outputStream := bufio.NewScanner(pipe)
	outputStream.Scan()

	line := outputStream.Text()
	root := line[strings.Index(line, "/"):]

	err = filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		// Exclude path containing offensive quotes
		if strings.Contains(path, "/off/") {
			return nil
		}
		// Filter out .dat files and directories
		if f.IsDir() || filepath.Ext(path) == ".dat" {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	i := randomInt(1, len(files))
	randomFile := files[i]

	// Open random file...
	file, err := os.Open(randomFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// ...and read its contents
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	quotes := string(content)

	quotesSlice := strings.Split(quotes, "%")
	q := randomInt(1, len(quotesSlice))

	fmt.Print(quotesSlice[q])
}
