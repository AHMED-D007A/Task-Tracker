package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func addition(file *os.File) {
	id := rand.Intn(900000) + 100000
	stringTask := fmt.Sprintf("\t{\"id\": %v, \"description\": \"%s\", \"status\": \"%s\", \"createdAt\": \"%v\", \"updatedAt\": \"%v\"},\n]", id, string(os.Args[2]), "todo", time.DateTime, time.DateTime)

	file.Seek(-2, io.SeekEnd)
	file.WriteString(stringTask)

	fmt.Printf("Task added successfully (ID: %v).\n", id)
}

func update(file *os.File) {
	scanner := bufio.NewScanner(file)
	var line string

	// temp file to replace the original file
	tempFile, err := os.Create("tempfile.json")
	if err != nil {
		log.Println(err)
		return
	}
	defer tempFile.Close()

	writer := bufio.NewWriter(tempFile)

	// reads until the task id is found and writes the rest in temp file.
	for scanner.Scan() {
		line = scanner.Text()
		if len(line) > 2 {
			if os.Args[2] == line[8:14] {
				break
			}
		}
		writer.WriteString(line + "\n")
	}
	// replace the description of the task and write the newline to tempfile.
	splitter := strings.Split(line, "\", \"status\":")
	newLine := line[:32] + os.Args[3] + "\", \"status\":" + splitter[1]

	writer.WriteString(newLine + "\n")

	// writes the rest of the original file in the temp file
	for scanner.Scan() {
		line = scanner.Text()
		writer.WriteString(line + "\n")
	}
	writer.Flush()

	// reSet the offset of the files to rewrite everything.
	tempFile.Seek(0, 0)
	file.Truncate(0)
	file.Seek(0, 0)

	writer = bufio.NewWriter(file)
	scanner = bufio.NewScanner(tempFile)
	for scanner.Scan() {
		line = scanner.Text()
		writer.WriteString(line + "\n")
	}
	writer.Flush()
	tempFile.Close()

	os.Remove("tempfile.json")
}

func deleteTask(file *os.File) {
	scanner := bufio.NewScanner(file)
	var line string

	// temp file to replace the original file
	tempFile, err := os.Create("tempfile.json")
	if err != nil {
		log.Println(err)
		return
	}
	defer tempFile.Close()

	writer := bufio.NewWriter(tempFile)

	// reads until the task id is found and writes the rest in temp file.
	for scanner.Scan() {
		line = scanner.Text()
		if len(line) > 2 {
			if os.Args[2] == line[8:14] {
				continue
			}
		}
		writer.WriteString(line + "\n")
	}
	writer.Flush()

	// reSet the offset of the files to rewrite everything.
	tempFile.Seek(0, 0)
	file.Truncate(0)
	file.Seek(0, 0)

	writer = bufio.NewWriter(file)
	scanner = bufio.NewScanner(tempFile)
	for scanner.Scan() {
		line = scanner.Text()
		writer.WriteString(line + "\n")
	}
	writer.Flush()
	tempFile.Close()

	os.Remove("tempfile.json")
}

func inProgress(file *os.File) {
	scanner := bufio.NewScanner(file)
	var line string

	// temp file to replace the original file
	tempFile, err := os.Create("tempfile.json")
	if err != nil {
		log.Println(err)
		return
	}
	defer tempFile.Close()

	writer := bufio.NewWriter(tempFile)

	// reads until the task id is found and writes the rest in temp file.
	for scanner.Scan() {
		line = scanner.Text()
		if len(line) > 2 {
			if os.Args[2] == line[8:14] {
				break
			}
		}
		writer.WriteString(line + "\n")
	}
	// replace the description of the task and write the newline to tempfile.
	splitter := strings.Split(line, "\", \"createdAt\":")
	splitter00 := strings.Split(splitter[0], "\", \"status\":")
	newLine := splitter00[0] + "\", \"status\":" + " \"in-progress" + "\", \"createdAt\":" + splitter[1]

	writer.WriteString(newLine + "\n")

	// writes the rest of the original file in the temp file
	for scanner.Scan() {
		line = scanner.Text()
		writer.WriteString(line + "\n")
	}
	writer.Flush()

	// reSet the offset of the files to rewrite everything.
	tempFile.Seek(0, 0)
	file.Truncate(0)
	file.Seek(0, 0)

	writer = bufio.NewWriter(file)
	scanner = bufio.NewScanner(tempFile)
	for scanner.Scan() {
		line = scanner.Text()
		writer.WriteString(line + "\n")
	}
	writer.Flush()
	tempFile.Close()

	os.Remove("tempfile.json")
}

func done(file *os.File) {
	scanner := bufio.NewScanner(file)
	var line string

	// temp file to replace the original file
	tempFile, err := os.Create("tempfile.json")
	if err != nil {
		log.Println(err)
		return
	}
	defer tempFile.Close()

	writer := bufio.NewWriter(tempFile)

	// reads until the task id is found and writes the rest in temp file.
	for scanner.Scan() {
		line = scanner.Text()
		if len(line) > 2 {
			if os.Args[2] == line[8:14] {
				break
			}
		}
		writer.WriteString(line + "\n")
	}
	// replace the description of the task and write the newline to tempfile.
	splitter := strings.Split(line, "\", \"createdAt\":")
	splitter00 := strings.Split(splitter[0], "\", \"status\":")
	newLine := splitter00[0] + "\", \"status\":" + " \"done" + "\", \"createdAt\":" + splitter[1]

	writer.WriteString(newLine + "\n")

	// writes the rest of the original file in the temp file
	for scanner.Scan() {
		line = scanner.Text()
		writer.WriteString(line + "\n")
	}
	writer.Flush()

	// reSet the offset of the files to rewrite everything.
	tempFile.Seek(0, 0)
	file.Truncate(0)
	file.Seek(0, 0)

	writer = bufio.NewWriter(file)
	scanner = bufio.NewScanner(tempFile)
	for scanner.Scan() {
		line = scanner.Text()
		writer.WriteString(line + "\n")
	}
	writer.Flush()
	tempFile.Close()

	os.Remove("tempfile.json")
}

func list(file *os.File) {
	switch os.Args[2] {
	case "done":
		scanner := bufio.NewScanner(file)
		var line string
		for scanner.Scan() {
			line = scanner.Text()
			if strings.Contains(line, "done") {
				fmt.Println(line[1:])
			}
		}
	case "todo":
		scanner := bufio.NewScanner(file)
		var line string
		for scanner.Scan() {
			line = scanner.Text()
			if strings.Contains(line, "todo") {
				fmt.Println(line[1:])
			}
		}
	case "in-progress":
		scanner := bufio.NewScanner(file)
		var line string
		for scanner.Scan() {
			line = scanner.Text()
			if strings.Contains(line, "in-progress") {
				fmt.Println(line[1:])
			}
		}
	}
}
