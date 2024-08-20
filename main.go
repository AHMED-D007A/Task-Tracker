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

func main() {
	var file *os.File

	// checks for file status.
	info, err := os.Stat("tasks.json")
	if os.IsNotExist(err) {
		file, err = os.Create("tasks.json")

		if err != nil {
			log.Println(err)
			return
		}

		file.Write([]byte("[\n]"))
	} else {
		file, err = os.OpenFile("tasks.json", os.O_RDWR, 0644)

		if err != nil {
			log.Println(err)
			return
		}

		if info.Size() == 0 {
			file.Write([]byte("[\n]"))
		}
	}

	if len(os.Args) == 1 {
		fmt.Println("Expected [command] [...argument].")
	} else {
		switch os.Args[1] {
		case "add":
			addition(file)
		case "update":
			update(file)
		case "delete":
			deleteTask(file)
		case "mark-in-progress":
		case "mark-done":
		case "list":
		}
	}
}

func addition(file *os.File) {
	id := rand.Intn(900000) + 100000
	stringTask := fmt.Sprintf("\t{\"id\": %v, \"description\": \"%s\", \"status\": \"%s\", \"createdAt\": \"%v\", \"updatedAt\": \"%v\"},\n]", id, string(os.Args[2]), "todo", time.DateTime, time.DateTime)

	file.Seek(-1, io.SeekEnd)
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
	file.Seek(0, 0)
	file.Truncate(0)

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
